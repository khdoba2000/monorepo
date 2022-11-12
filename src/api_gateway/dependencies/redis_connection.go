package dependencies

import (
	"monorepo/src/api_gateway/configs"
	"monorepo/src/libs/redis"
	"time"

	rd "github.com/gomodule/redigo/redis"
)

func NewRedisConn() redis.InMemoryStorageI {
	pl := &rd.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (rd.Conn, error) { return rd.Dial("tcp", configs.Config().RedisPort) },
	}

	rd := redis.NewRedisRepo(pl)

	return rd
}
