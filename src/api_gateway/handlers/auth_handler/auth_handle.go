package auth_handler

import (
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/idl/auth_service"
	"monorepo/src/libs/log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	//"monorepo/src/api_gateway/ci"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	logger     log.Factory
	authClient auth_service.AuthServiceClient
}

// New creates auth handlers
func New(logger log.Factory) AuthHandlers {

	return &authHandler{
		logger:     logger,
		authClient: dependencies.AuthServiceClient(),
	}
}

func (ah *authHandler) TestHandler(w http.ResponseWriter, r *http.Request) {

	ah.logger.For(r.Context()).Info("TestHandler hit")

	ah.authClient.StaffSignUp(opentracing.ContextWithSpan(r.Context(), opentracing.SpanFromContext(r.Context())), &auth_service.CustomerSignUpRequest{Name: "name"})

	ah.logger.For(r.Context()).Info("TestHandler success")

	w.Write([]byte("Hello, World1!"))
}

func (ah *authHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	// container := ci.Get()
}
