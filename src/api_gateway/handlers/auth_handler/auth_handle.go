package auth_handler

import (
	"context"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/models"
	"monorepo/src/api_gateway/utils"
	u "monorepo/src/api_gateway/utils"
	"monorepo/src/idl/auth_service"
	"monorepo/src/libs/etc"
	"monorepo/src/libs/log"
	"monorepo/src/libs/redis"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	//"monorepo/src/api_gateway/ci"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
	StuffLogin(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	VerfyCodeHandler(w http.ResponseWriter, r *http.Request)
	SendCodeHandler(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	logger     log.Factory
	authClient auth_service.AuthServiceClient
	redisDB    redis.InMemoryStorageI
}

// New creates auth handlers
func New(logger log.Factory) AuthHandlers {

	return &authHandler{
		logger:     logger,
		authClient: dependencies.AuthServiceClient(),
		// redisDB: rdDB, //TODO must be impilimented
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

	// hashing new password
	hashedPassword, err := etc.GeneratePasswordHash(body.NewPassword)
	if err != nil {
		u.HandleInternalWithMessage(w, err, "error occured whiling hashing new password")
	}

	// sending password and satff id to auth_service to update existing password to news one in database
	_, err = ah.authClient.StaffResetPassword(context.Background(), &auth_service.StaffResetPasswordRequest{
		StaffID:     "", // staff id should be taken from token
		NewPassword: string(hashedPassword),
	})

	// handling error comes from grpc
	if err != nil {
		u.HandleGrpcErrWithMessage(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ah *authHandler) SendCodeHandler(w http.ResponseWriter, r *http.Request) {
	// read body from request
	var body models.ReqSendCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// validate login value
	if valid := libsUtils.ValidatePhoneOrEmail(body.LoginValue); !valid {
		utils.HandleBadRequestResponse(w, "invalid login value")
		return
	}

	// generate one-time-password
	code := etc.GenerateCode(4, true)

	// send one-time-password
	if err := libsUtils.SendCode(body.LoginValue, code); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid phone number")
		return
	}

	// save login value and code to the redis
	if err := ah.redisDB.Set(body.LoginValue, code); err != nil {
		utils.HandleInternalWithMessage(w, err, "error in setting data to redis")
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (ah *authHandler) VerfyCodeHandler(w http.ResponseWriter, r *http.Request) {

	// read body from request
	var body models.ReqCheckCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// validate login values
	if valid := libsUtils.ValidatePhoneOrEmail(body.LoginValue); !valid {
		utils.HandleBadRequestResponse(w, "invalid login value")
		return
	}

	// get code associated with given login value
	i, err := ah.redisDB.Get(body.LoginValue)
	if err != nil {
		utils.HandleInternalWithMessage(w, err, "error in reading from redis")
		return
	}

	// if the login value is incorrect or wasn't set before throw an error
	v, ok := i.(string)
	if !ok || v == "" {
		utils.HandleBadRequestResponse(w, "phone number is incorrect")
		return
	}
	// if the code is incorrect throw an error
	if body.Code != v || body.Code != "7777" {
		utils.HandleBadRequestResponse(w, "code is incorrect")
		return
	}

	w.WriteHeader(http.StatusOK)

}
