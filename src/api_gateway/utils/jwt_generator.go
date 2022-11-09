package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"monorepo/src/api_gateway/configs"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var conf = configs.Config()

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id string, credentials map[string]string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(id, credentials)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id string, credentials map[string]string) (string, error) {

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(conf.JWTRefreshKeyExpireHours)).Unix()
	claims["role"] = credentials["role"]
	//	// Set private token credentials:
	//	for _, credential := range credentials {
	//		claims[credential] = true
	//	}

	// in local server access token ttl = 10 days
	if conf.Environment == "develop" {
		claims["expires"] = time.Now().Add(time.Minute * time.Duration(10*conf.JWTRefreshKeyExpireHours)).Unix()
	} else {
		// in staging server access token ttl = day
		claims["expires"] = time.Now().Add(time.Minute * time.Duration(conf.JWTRefreshKeyExpireHours)).Unix()
	}
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(conf.JWTSecretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	sha256 := sha256.New()

	// Create a new now date and time string with salt.
	refresh := conf.JWTSecretKey + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := sha256.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(conf.JWTRefreshKeyExpireHours)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(sha256.Sum(nil)) + "." + expireTime

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
