package redis

import (
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/log"
	"sync"
)

type RedisManager struct {
	redisPool RedisPool
}

var m *RedisManager
var once sync.Once

func GetInstance() *RedisManager {
	once.Do(func() {
		m = &RedisManager{}
	})
	return m
}

func (self *RedisManager) InitManager() error {
	host, err := config.GetInstance().GetConfig("REDIS_ADDR")
	if nil != err {
		log.Error("RedisMAnager InitManager fail! Can not find conf REDIS_HOST!!!")
		return err
	}
	passwd, err := config.GetInstance().GetConfig("REDIS_PASSWD")
	if nil != err {
		log.Error("RedisMAnager InitManager fail! Can not find conf REDIS_PASSWD!!!")
		return err
	}
	self.redisPool.Init(host, passwd)
	return nil
}

func (self *RedisManager) SetStringValue(key, value string) error {
	return self.redisPool.SetStringValue(key, value)
}

func (self *RedisManager) SetStringValueWithExpireTime(key, value string, expireTime int64) error {
	return self.redisPool.SetStringValueWithExpireTime(key, value, expireTime)
}

func (self *RedisManager) GetStringValue(key string) (string, error) {
	return self.redisPool.GetStringValue(key)
}
