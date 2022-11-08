package configs

import (
	"errors"
	"sync"
	"time"

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

const (
	// NewStatus enum
	NewStatus = "new"
	// SentStatus enum
	SentStatus = "sent"
	// SubscribeType enum
	SubscribeType = "subscribe"
	// PostType enum
	PostType = "post"
	// UnsubscribeType enum
	UnsubscribeType = "unsubscribe"
	// LikeType enum
	LikeType = "like"
	// JoinType enum
	JoinType = "join"
	// CommentType enum
	CommentType = "comment"
	// DislikeType enum
	DislikeType = "dislike"
	// Replied enum
	Replied = "replied"
	// TipType enum
	TipType = "tip"
	// PurchaseType enum
	PurchaseType = "purchase"
	// Email enum
	Email = "email"
)

// Configuration ...
type Configuration struct {
	HTTPPort    string `json:"http_port"`
	LogLevel    string `json:"log_level"`
	Environment string `json:"environment"`

	BucketName                 string
	StorageType                string
	AwsS3Id                    string
	AwsS3Secret                string
	AwsBucketURL               string
	AppURL                     string
	PostgresHost               string
	PostgresPort               int
	PostgresDatabase           string
	PostgresUser               string
	PostgresPassword           string
	ServerPort                 int
	ServerHost                 string
	ServiceDir                 string
	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration

	RedisHost string
	RedisPort int

	CasbinConfigPath    string
	MiddlewareRolesPath string

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	FirebaseWebKey     string
	DomainURIPrefix    string
	AndroidPackageName string
	IosBundleID        string

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int

	SendgridAPIKey     string
	SendgridEmail      string
	SMTPUser           string
	SMTPUserPass       string
	SMTPHost           string
	SMTPPort           int
	EmailFromHeader    string
	SocialGoogleKey    string
	SocialGoogleSecret string

	SocialFacebookKey    string
	SocialFacebookSecret string
	TwilioFrom           string
	TwilioAccountSid     string
	TwilioAuthToken      string
}

func load() *Configuration {

	var config Configuration

	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("BUCKET_NAME", "medley-storage")
	v.SetDefault("ENVIRONMENT", "development")
	v.SetDefault("LOG_LEVEL", "debug")
	v.SetDefault("HTTP_PORT", ":8000")
	v.SetDefault("CASBIN_CONFIG_PATH", "src/api_gateway/configs/rbac_model.conf")
	v.SetDefault("MIDDLEWARE_ROLES_PATH", "src/api_gateway/configs/models.csv")
	v.SetDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)
	v.SetDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)
	v.SetDefault("POSTGRES_HOST", "localhost")
	v.SetDefault("POSTGRES_PORT", 5432)
	v.SetDefault("POSTGRES_USER", "muhammad")
	v.SetDefault("POSTGRES_PASSWORD", "12345")
	v.SetDefault("POSTGRES_DB", "mono_db")
	v.SetDefault("REDIS_HOST", "localhost")
	v.SetDefault("REDIS_PORT", 6379)
	v.SetDefault("CONTEXT_TIMEOUT", 7)

	config.Environment = v.GetString("ENVIRONMENT")
	config.BucketName = v.GetString("BUCKET_NAME")
	config.AwsS3Id = v.GetString("AWS_S3_ID")
	config.AwsS3Secret = v.GetString("AWS_S3_SECRET")
	config.HTTPPort = v.GetString("HTTP_PORT")
	config.LogLevel = v.GetString("LOG_LEVEL")
	config.CasbinConfigPath = v.GetString("CASBIN_CONFIG_PATH")
	config.MiddlewareRolesPath = v.GetString("MIDDLEWARE_ROLES_PATH")
	config.JWTSecretKey = v.GetString("JWT_SECRET_KEY")
	config.JWTRefreshKey = v.GetString("JWT_REFRESH_KEY")
	config.JWTSecretKeyExpireMinutes = v.GetInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	config.JWTRefreshKeyExpireHours = v.GetInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	config.PostgresDatabase = v.GetString("POSTGRES_DB")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresPort = v.GetInt("POSTGRES_PORT")
	config.RedisHost = v.GetString("REDIS_HOST")
	config.RedisPort = v.GetInt("REDIS_PORT")
	config.SendgridEmail = v.GetString("SENDGRID_EMAIL")
	config.SendgridAPIKey = v.GetString("SENDGRID_API_KEY")
	config.CtxTimeout = v.GetInt("CONTEXT_TIMEOUT")
	return &config
}

// Validate validates the configuration
func (c *Configuration) Validate() error {
	if c.HTTPPort == "" {
		return errors.New("http_port required")
	}
	return nil
}
