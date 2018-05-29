package redis

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/log"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	REDIS_POOL_MAX_IDLE     = 512
	REDIS_POOL_MAX_ACTIVE   = 1024
	REDIS_POOL_IDLE_TIMEOUT = 240
)

type RedisPool struct {
	redisPool *redis.Pool
	isInit    bool
}

func (self *RedisPool) Init(host, passwd string) {
	self.redisPool = self.newPool(host, passwd)
	self.isInit = true
}

func (self *RedisPool) newPool(host, passwd string) *redis.Pool {
	maxIdle, err := config.GetInstance().GetInt("REDIS_POOL_MAX_IDLE")
	if nil != err {
		log.Error(err.Error())
		maxIdle = REDIS_POOL_MAX_IDLE
	}
	maxActive, err := config.GetInstance().GetInt("REDIS_POOL_MAX_ACTIVE")
	if nil != err {
		log.Error(err.Error())
		maxActive = REDIS_POOL_MAX_ACTIVE
	}
	idleTimeout, err := config.GetInstance().GetInt("REDIS_POOL_IDLE_TIMEOUT")
	if nil != err {
		log.Error(err.Error())
		idleTimeout = REDIS_POOL_IDLE_TIMEOUT
	}
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if "" != passwd {
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (self *RedisPool) setValue(key string, value interface{}) error {
	if !self.isInit {
		return errors.New("RedisPool setValue fail!,Has not Init")
	}
	conn := self.redisPool.Get()
	if nil == conn {
		return errors.New("RedisPool setValue fail!,Can not get conn!!!")
	}
	defer conn.Close()
	ret, err := conn.Do("SET", key, value)
	if nil != err {
		return err
	}
	if ret != "OK" {
		return errors.New("setValue fail! do set return err!!!")
	}
	return err
}

func (self *RedisPool) setValueWithExpireTime(key string, value interface{}, expireTime int64) error {
	if !self.isInit {
		return errors.New("RedisPool setValue fail!,Has not Init")
	}
	conn := self.redisPool.Get()
	if nil == conn {
		return errors.New("RedisPool setValue fail!,Can not get conn!!!")
	}
	defer conn.Close()
	ret, err := conn.Do("SET", key, value)
	if nil != err {
		return err
	}
	if ret != "OK" {
		return errors.New("setValueWithExpireTime fail! do set return err!!!")
	}
	ret, err = conn.Do("EXPIRE", key, expireTime)
	if nil != err {
		return err
	}
	if ret != int64(1) {
		return errors.New("setValueWithExpireTime fail! do expire return err!!!")
	}
	return err
}

func (self *RedisPool) SetStringValue(key, value string) error {
	err := self.setValue(key, value)
	if nil != err {
		return err
	}
	return err
}

func (self *RedisPool) SetStringValueWithExpireTime(key, value string, expireTime int64) error {
	err := self.setValueWithExpireTime(key, value, expireTime)
	if nil != err {
		return err
	}
	return err
}

func (self *RedisPool) getValue(key string) (interface{}, error) {
	if !self.isInit {
		return "", errors.New("RedisPool getValue fail!,Client Unconnect!!!")
	}
	conn := self.redisPool.Get()
	if nil == conn {
		return "", errors.New("RedisPool getValue fail!,Can not get conn!!!")
	}
	defer conn.Close()
	return conn.Do("GET", key)
}

func (self *RedisPool) GetStringValue(key string) (string, error) {
	return redis.String(self.getValue(key))
}

func (self *RedisPool) setKeyExpireTime(key string, expireTime int64) error {
	if !self.isInit {
		return errors.New("RedisPool setKeyExpireTime fail!,Client Unconnect!!!")
	}
	conn := self.redisPool.Get()
	if nil == conn {
		return errors.New("RedisPool setKeyExpireTime fail!,Can not get conn!!!")
	}
	defer conn.Close()
	ret, err := conn.Do("EXPIRE", key, expireTime)
	if nil != err {
		return err
	}
	if ret != "OK" {
		return errors.New("setKeyExpireTime fail! do expire return err!!!")
	}
	return err
}
