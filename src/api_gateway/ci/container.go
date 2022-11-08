package ci

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gomodule/redigo/redis"

	"monorepo/src/api_gateway/configs"
	repo "monorepo/src/libs/redis"
)

// Container ...
type Container struct {
	Redis repo.InMemoryStorageI
	S3    *s3.S3
}

var (
	instance *Container
	once     sync.Once

	cfg = configs.Config()
)

// Get - get Container instance
func Get() *Container {
	once.Do(func() {
		sess, err := session.NewSession(&aws.Config{
			// TODO Region should be moved to config
			Region:      aws.String("ap-south-1"),
			Credentials: credentials.NewStaticCredentials(cfg.AwsS3Id, cfg.AwsS3Secret, ""),
		})

		if err != nil {
			panic(err)
		}
		instance = &Container{
			Redis: redisPool(),
			S3:    s3.New(sess),
		}
	})

	return instance
}

func redisPool() repo.InMemoryStorageI {
	pool := redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	return repo.NewRedisRepo(&pool)
}
