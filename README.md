# Simple Auth
## Version 1.0.0

# HTTP API

## Auth API
### http://service_host:port/auth
### 使用方法
#### 参数
##### Request 
###### 参数1：ClientID     类型：string
###### 参数2：ClientSecret 类型：string
##### Response
###### 参数1：Result       类型：string
###### 参数2：Token        类型：string

## Verification API
### http://service_host:port/verification
### 使用方法
#### 参数
##### Request 
###### 参数1：Token        类型：string
##### Response
###### 参数1：Result       类型：string
###### 参数2：ClientID     类型：string

## Registry API
### http://service_host:port/registry
### 使用方法
#### 参数
##### Request 
###### 参数1：ClientID     类型：string
###### 参数2：ClientSecret 类型：string
##### Response
###### 参数1：Result       类型：string

## GRPC API
#### Proto 见 github.com/seaofstars-coder/simple-auth/proto包中auth.proto和verification.proto


# ASR-AUTH 配置说明
~~~~~~
# HTTP 服务监听端口
AUTH_HTTP_ADDR::9999
# Redis 服务地址端口
REDIS_ADDR:47.52.227.23:6379
# Redis 服务密码
REDIS_PASSWD:
# RedisPool 参数 最大空闲连接数量
REDIS_POOL_MAX_IDLE:512
# RedisPool 参数 最大活跃连接数量
REDIS_POOL_MAX_ACTIVE:1024
# RedisPool 参数 空闲连接超时时间
REDIS_POOL_IDLE_TIMEOUT:240
# 鉴权Token有效时间
AUTH_TOKEN_EXPIRE_TIME:60
# GRPC 服务监听端口
GRPC_ADDR::19999
# Zookeeper 服务地址端口
ZOOKEEPER_ADDRS:47.52.227.23:2181
# Zookeeper 连接超时时间
ZOOKEEPER_TIMEOUT:60
# Zookeeper 客户数据路径
ZOOKEEPER_CLIENTS_PATH:/ovs/clients
~~~~~~
