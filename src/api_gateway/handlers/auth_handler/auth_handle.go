package auth_handler

import (
	"context"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/models"
	u "monorepo/src/api_gateway/utils"
	"monorepo/src/idl/auth_service"
	"monorepo/src/libs/log"
	"net/http"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	//"monorepo/src/api_gateway/ci"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
	StuffLogin(w http.ResponseWriter, r *http.Request)
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

func (ah *authHandler) StuffLogin(w http.ResponseWriter, r *http.Request) {
	var (
		body models.StuffLoginModel
	)

	err := u.BodyParser(r, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.WriteJSON(w, err.Error())
		return
	}

	if !strings.Contains(body.PhoneNumber, "+") || len(body.PhoneNumber) != 13 {
		w.WriteHeader(http.StatusBadRequest)
		u.WriteJSON(w, "error: phone number is not correctly filled")
		return
	}

	err = body.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.WriteJSON(w, "error: one of the fields is not correct")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Second*time.Duration(configs.Config().CtxTimeout))
	defer cancel()

	_, err = dependencies.AuthServiceClient().StaffLogin(ctx, &auth_service.StaffLoginRequest{
		PhoneNumber: body.PhoneNumber,
		Password:    body.Password,
	})
	if err != nil {
		u.HandleGrpcErrWithMessage(w, err, "Error in Login")
		return
	}

}
