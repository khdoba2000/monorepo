package auth_handler

import (
	"monorepo/src/libs/logger"
	"net/http"

	//"monorepo/src/api_gateway/ci"

	"github.com/opentracing/opentracing-go"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	tracer opentracing.Tracer
}

// New creates auth handlers
func New(logger logger.Logger) AuthHandlers {
	return &authHandler{}
}

func (ah *authHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World1!"))
}

func (ah *authHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	// container := ci.Get()
}
