package redis

import (
	rd "github.com/gomodule/redigo/redis"
)

// InMemoryStorageI is interface for storage in memory
type InMemoryStorageI interface {
	Set(key, value string) error
	SetWithTTL(key, value string, seconds int) error
	Get(key string) (interface{}, error)
}

type redisRepo struct {
	rds *rd.Pool
}

// NewRedisRepo ...
func NewRedisRepo(rds *rd.Pool) InMemoryStorageI {
	return &redisRepo{rds: rds}
}

func (r *redisRepo) Set(key, value string) (err error) {
	conn := r.rds.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	return
}

// SetWithTTL ...
func (r *redisRepo) SetWithTTL(key, value string, seconds int) (err error) {
	conn := r.rds.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, seconds, value)
	return
}

func (r *redisRepo) Get(key string) (interface{}, error) {
	conn := r.rds.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}
