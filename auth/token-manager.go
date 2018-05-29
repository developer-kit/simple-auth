package auth

import (
	"encoding/json"
	"fmt"
	"github.com/seaofstars-coder/simple-auth/common"
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/log"
	"github.com/seaofstars-coder/simple-auth/redis"
	"github.com/seaofstars-coder/simple-auth/util"
	"strings"
)

type AuthTokenManager struct {
	clientTokenMap ClientAuthTokenMap
}

func (self *AuthTokenManager) Init() {
	self.clientTokenMap = make(ClientAuthTokenMap)
}

func (self *AuthTokenManager) GenerateAuthToken(clientID, clientSecret string) string {
	token := util.NewUUID()
	token = util.StringFixedLen(token, 32)
	f_clientID := util.StringFixedLen(clientID, 64)
	f_clientSecret := util.StringFixedLen(clientSecret, 64)
	token = token + f_clientID + f_clientSecret
	token, _ = util.Hash("sha512", token)
	return strings.ToUpper(fmt.Sprintf("%x", token))
}

func (self *AuthTokenManager) GetAuthToken(clientID, clientSecret string) string {
	value, err := redis.GetInstance().GetStringValue(self.GenerateAuthClientKey(clientID, clientSecret))
	if nil != err {
		return self.NewAuthToken(clientID, clientSecret)
	}
	return value
}

func (self *AuthTokenManager) NewAuthToken(clientID, clientSecret string) string {
	token := self.GenerateAuthToken(clientID, clientSecret)
	expireTime, err := config.GetInstance().GetInt("AUTH_TOKEN_EXPIRE_TIME")
	if nil != err {
		log.Error(err.Error())
		return ""
	}
	err = redis.GetInstance().SetStringValueWithExpireTime(self.GenerateAuthTokenKey(token), self.GenerateAuthData(clientID, clientSecret), int64(expireTime))
	if nil != err {
		log.Error(err.Error())
		return ""
	}
	err = redis.GetInstance().SetStringValueWithExpireTime(self.GenerateAuthClientKey(clientID, clientSecret), token, int64(expireTime))
	if nil != err {
		log.Error(err.Error())
		return ""
	}
	log.Info("NewAuthToken , ClientId:", clientID, "\tClientSecret:", clientSecret, "\tToken:", token)
	return token
}

func (self *AuthTokenManager) VerificationAuthToken(token string) (bool, string) {
	data, err := redis.GetInstance().GetStringValue(self.GenerateAuthTokenKey(token))
	if nil != err {
		log.Error(err.Error())
		return false, ""
	}
	var clientAuthData common.ClientAuthData
	json.Unmarshal([]byte(data), &clientAuthData)
	return true, clientAuthData.ClientId
}

func (self AuthTokenManager) GenerateAuthTokenKey(token string) string {
	return "client_auth_token:" + token
}

func (self AuthTokenManager) GenerateAuthClientKey(clientID, clientSecret string) string {
	return "client_auth_client:" + clientID + "@" + clientSecret
}

func (self AuthTokenManager) GenerateAuthData(clientID, clientSecret string) string {
	clientAuthData := common.ClientAuthData{ClientId: clientID, ClientSecret: clientSecret}
	data, err := json.Marshal(clientAuthData)
	if nil != err {
		return ""
	}
	return string(data)
}
