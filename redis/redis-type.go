package redis

type RedisClientStatus uint

const (
	RCS_CONNECTED  = RedisClientStatus(1)
	RCS_CONNECTING = RedisClientStatus(2)
	RCS_UNCONNECT  = RedisClientStatus(3)
)
