package configs

import (
	"errors"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	conf *Configuration
	once sync.Once
)

// Config loads configuration using atomic pattern
func Config() *Configuration {
	once.Do(func() {
		conf = load()
	})
	return conf
}

// Configuration ...
type Configuration struct {
	HTTPPort    string `json:"http_port"`
	LogLevel    string `json:"log_level"`
	Environment string `json:"environment"`

	ServerPort                 int
	ServerHost                 string
	ServiceDir                 string
	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration

	RedisHost       string
	RedisPort       string
	AuthServiceHost string
	AuthServicePort int

	CasbinConfigPath    string
	MiddlewareRolesPath string

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
}

func load() *Configuration {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load("src/api_gateway/.env")
	if err != nil {
		panic(err)
	}

	var config Configuration

	v := viper.New()
	v.AutomaticEnv()

	config.Environment = v.GetString("ENVIRONMENT")
	config.HTTPPort = v.GetString("HTTP_PORT")
	config.LogLevel = v.GetString("LOG_LEVEL")
	config.CasbinConfigPath = v.GetString("CASBIN_CONFIG_PATH")
	config.MiddlewareRolesPath = v.GetString("MIDDLEWARE_ROLES_PATH")
	config.JWTSecretKey = v.GetString("JWT_SECRET_KEY")
	config.JWTRefreshKey = v.GetString("JWT_REFRESH_KEY")
	config.JWTSecretKeyExpireMinutes = v.GetInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	config.JWTRefreshKeyExpireHours = v.GetInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	config.RedisHost = v.GetString("REDIS_HOST")
	config.RedisPort = v.GetString("REDIS_PORT")
	config.AuthServiceHost = v.GetString("AUTH_SERVICE_HOST")
	config.AuthServicePort = v.GetInt("AUTH_SERVICE_PORT")
	config.CtxTimeout = v.GetInt("CONTEXT_TIMEOUT")
	return &config
}

func (c *Configuration) validate() error {
	if c.HTTPPort == "" {
		return errors.New("http_port required")
	}
	return nil
}
