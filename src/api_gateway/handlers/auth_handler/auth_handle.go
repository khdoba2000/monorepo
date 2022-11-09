package auth_handler

import (
	"monorepo/src/api_gateway/pkg/log"
	"net/http"
	//"monorepo/src/api_gateway/ci"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	logger log.Factory
}

// New creates auth handlers
func New(logger log.Factory) AuthHandlers {
	return &authHandler{
		logger: logger,
	}
}

func (ah *authHandler) TestHandler(w http.ResponseWriter, r *http.Request) {

	ah.logger.For(r.Context()).Info("TestHandler hit")
	w.Write([]byte("Hello, World1!"))
}

func (ah *authHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	// container := ci.Get()
}
