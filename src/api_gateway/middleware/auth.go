package middleware

import (
	"log"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/libs/jwt"
	"net/http"

	"github.com/casbin/casbin/v2"
)

// JWTRoleAuthorizer is a sturcture for a Role Authorizer type
type JWTRoleAuthorizer struct {
	enforcer interface {
		Enforce(rvals ...interface{}) (bool, error)
	}
	signingKey []byte
	//	logger     logger.Logger
}

// NewCasbinJWTRoleAuthorizer creates and returns new Role Authorizer
func NewCasbinJWTRoleAuthorizer(cfg *configs.Configuration) (*JWTRoleAuthorizer, error) {

	enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
	if err != nil {
		log.Println("could not initialize new enforcer:", err.Error())
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		signingKey: []byte(cfg.JWTSecretKey),
		//		logger:     logger,
	}, nil
}

// Middleware ...
func (jwta *JWTRoleAuthorizer) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check Permission with casbin
		allowed, err := jwta.checkPermission(r)
		if err != nil {
			// Casbin.Enforcer not working normal
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		if !allowed {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})
}

func (jwta *JWTRoleAuthorizer) checkPermission(r *http.Request) (bool, error) {

	role, err := jwta.getUser(r.Header.Get("Authorization"))
	if err != nil {
		return false, err
	}

	method := r.Method
	path := r.URL.Path
	enforsed, err := jwta.enforcer.Enforce(role, path, method)
	return enforsed, err
}

func (jwta *JWTRoleAuthorizer) getUser(accessToken string) (string, error) {

	claims, err := jwt.ExtractClaims(accessToken, jwta.signingKey)
	if err != nil {
		log.Println("could not extract claims:", err)
		return "", err
	}
	return claims["role"].(string), nil
}
