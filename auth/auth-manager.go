package auth

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/redis"
	"github.com/seaofstars-coder/simple-auth/zookeeper"
	"sync"
)

type AuthManager struct {
	handlerFuncMap AuthHandlerFuncMap
	tokenManager   AuthTokenManager
}

var m *AuthManager
var once sync.Once

func GetInstance() *AuthManager {
	once.Do(func() {
		m = &AuthManager{}
	})
	return m
}

func (self *AuthManager) InitManager() error {
	self.handlerFuncMap = make(AuthHandlerFuncMap)
	self.initHandlerFuncMap()
	return nil
}

func (self *AuthManager) GetAuthTokenManager() *AuthTokenManager {
	return &self.tokenManager
}

func (self *AuthManager) Handle(request AuthServiceRequest) AuthServiceResponse {
	handleFunc := self.handlerFuncMap[request.ServiceType]
	if nil == handleFunc {
		return AuthServiceResponse{}
	}
	rsp, err := handleFunc(request.Req)
	if nil != err {
		return AuthServiceResponse{}
	}
	rsp.ServiceType = request.ServiceType
	return rsp
}

func (self *AuthManager) initHandlerFuncMap() {
	self.RegisterHandler(AST_AUTH, handleAuthRequest)
	self.RegisterHandler(AST_VERIFICATION, handleVerificationRequest)
	self.RegisterHandler(AST_REGISTRY, handleRegistryRequest)
}

func (self *AuthManager) RegisterHandler(sT AuthServiceType, hFunc AuthHandlerFunc) error {
	if self.handlerFuncMap[sT] != nil {
		return errors.New("RegisterHandler Fail! Has Func Registered Before!!!")
	}
	self.handlerFuncMap[sT] = hFunc
	return nil
}

func (self *AuthManager) UnregisterHandler(sT AuthServiceType) {
	self.handlerFuncMap[sT] = nil
}

func (self *AuthManager) GetTokenValidTime() int64 {
	return AuthTokenValidTime
}

func (self AuthManager) CheckClientSecret(clientID, clientSecret string) bool {
	secret, err := zookeeper.GetInstance().GetClientSecret(clientID)
	if nil != err {
		secret, err = redis.GetInstance().GetStringValue(self.GenerateClientKey(clientID))
		if nil != err {
			return false
		}
		return clientSecret == secret
	}
	return clientSecret == secret
}

func (self AuthManager) RegistryClient(clientID, clientSecret string) error {
	return redis.GetInstance().SetStringValue(self.GenerateClientKey(clientID), clientSecret)
}

func (self AuthManager) GenerateClientKey(clientID string) string {
	return "client_account:" + clientID
}
