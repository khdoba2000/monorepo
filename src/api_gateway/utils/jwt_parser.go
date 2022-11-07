package utils

import (
	"log"
	"monorepo/src/libs/jwt"

	// "github.com/dgrijalva/jwt-go"

	"github.com/google/uuid"
)

// GetUserIDFromToken ...
func GetUserIDFromToken(accessToken string) (uuid.UUID, error) {

	claims, err := jwt.ExtractClaims(accessToken, []byte(conf.JWTSecretKey))
	if err != nil {
		log.Println("could not extract claims:", err)
		return uuid.Nil, err
	}
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID      uuid.UUID
	Credentials map[string]string
	Expires     int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
// func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
// 	token, err := verifyToken(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Setting and checking token and credentials.
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		// User ID.
// 		userID, err := uuid.Parse(claims["id"].(string))
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Expires time.
// 		expires := int64(claims["expires"].(float64))

// 		// User credentials.
// 		credentials := map[string]string{
// 			"role": claims["role"].(string),
// 		}

// 		return &TokenMetadata{
// 			UserID:      userID,
// 			Credentials: credentials,
// 			Expires:     expires,
// 		}, nil
// 	}

// 	return nil, err
// }

// func extractToken(c *fiber.Ctx) string {
// 	bearToken := c.Get("Authorization")

// 	onlyToken := strings.Split(bearToken, " ")
// 	if len(onlyToken) == 1 {
// 		return onlyToken[0]
// 	}

// 	return ""
// }

// func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
// 	tokenString := extractToken(c)
// 	token, err := jwt.Parse(tokenString, jwtKeyFunc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
// 	return []byte(conf.JWTSecretKey), nil
// }
