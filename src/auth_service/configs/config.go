package configs

import (
	"errors"
	"sync"

	"github.com/joho/godotenv"

	"github.com/spf13/viper"
)

var (
	instance *Configuration
	once     sync.Once
)

// Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	LogLevel    string `json:"log_level"`
	Environment string `json:"environment"`

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	ServerPort       int
	ServerHost       string
	ServiceDir       string

	RPCPort string

	// context timeout in seconds
	CtxTimeout        int
	ServerReadTimeout int
}

func load() *Configuration {

	// load .env file from given path
	err := godotenv.Load("src/auth_service/.env")
	if err != nil {
		panic(err)
	}

	var config Configuration

	v := viper.New()
	v.AutomaticEnv()

	config.Environment = v.GetString("ENVIRONMENT")
	config.LogLevel = v.GetString("LOG_LEVEL")
	config.PostgresDatabase = v.GetString("POSTGRES_DB")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresPort = v.GetInt("POSTGRES_PORT")
	config.RPCPort = v.GetString("RPC_PORT")
	config.CtxTimeout = v.GetInt("CONTEXT_TIMEOUT")
	return &config
}

// Validate validates the configuration
func (c *Configuration) Validate() error {
	if c.RPCPort == "" {
		return errors.New("rpc_port required")
	}
	return nil
}
