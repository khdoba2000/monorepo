package auth_handler

import (
	"context"
	"fmt"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/models"
	"monorepo/src/api_gateway/utils"
	u "monorepo/src/api_gateway/utils"
	"monorepo/src/idl/auth_service"
	"monorepo/src/libs/etc"
	"monorepo/src/libs/log"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
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

	ah.authClient.StaffSignUp(opentracing.ContextWithSpan(r.Context(), opentracing.SpanFromContext(r.Context())), &auth_service.StaffSignUpRequest{Name: "name"})

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

	err = body.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.WriteJSON(w, "error: one of the fields is not correct")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*time.Duration(configs.Config().CtxTimeout))
	defer cancel()

	res, err := dependencies.AuthServiceClient().StaffLogin(ctx, &auth_service.StaffLoginRequest{
		PhoneNumber: body.PhoneNumber,
		Password:    body.Password,
	})
	if err != nil {
		u.HandleGrpcErrWithMessage(w, err, "Error in Login")
		return
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(res.Id, map[string]string{
		"role": res.Role,
	}, res.BranchId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.WriteJSON(w, err.Error())
		return
	}

	u.WriteJSON(w, models.StaffLoginResponse{
		ID:           res.Id,
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	})

}

func (ah *authHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// read body from request
	var body models.ReqResetPassword
	if err := u.BodyParser(r, body); err != nil {
		u.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// check if the password meets requirements
	if err := libsUtils.ValidatePassword(body.NewPassword); err != nil {
		u.HandleBadRequestResponse(w, err.Error())
		return
	}

	hashedPassword, err := etc.GeneratePasswordHash(body.NewPassword)
	if err != nil {
		u.HandleInternalWithMessage(w, err, "error occured whiling hashing new password")
	}

	// this is just to avoid unused variable error. Later we will send it to service where this hashed password will be saved to database
	fmt.Println(hashedPassword)
	// here we should update password from databse

	w.WriteHeader(http.StatusOK)
}
