package customer_handler

import (
	"monorepo/src/api_gateway/models"
	"monorepo/src/api_gateway/utils"
	"monorepo/src/libs/etc"
	"monorepo/src/libs/logger"
	"monorepo/src/libs/redis"
	libs "monorepo/src/libs/utils"
	"net/http"
)

type CustomerHandlers interface {
	TestHandler2(w http.ResponseWriter, r *http.Request)
	VerfyCodeHandler(w http.ResponseWriter, r *http.Request)
	SendCodeHandler(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	redisDB redis.InMemoryStorageI
}

// New creates handler
func New(logger logger.Logger) CustomerHandlers {
	return &customerHandler{}
}
func (ch *customerHandler) TestHandler2(w http.ResponseWriter, r *http.Request) {
	// ch.logger.Print("Got a request.")
	w.Write([]byte("Hello, World2!"))
}

func (ch *customerHandler) SendCodeHandler(w http.ResponseWriter, r *http.Request) {
	// read body from request
	var body models.ReqSendCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// validate login value
	if valid := libs.ValidatePhoneOrEmail(body.LoginValue); !valid {
		utils.HandleBadRequestResponse(w, "invalid login value")
		return
	}

	// generate one-time-password
	code := etc.GenerateCode(4, true)

	// send one-time-password
	if err := libs.SendCode(body.LoginValue, code); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid phone number")
		return
	}

	// save login value and code to the redis
	if err := ch.redisDB.Set(body.LoginValue, code); err != nil {
		utils.HandleInternalWithMessage(w, err, "error in setting data to redis")
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (ch *customerHandler) VerfyCodeHandler(w http.ResponseWriter, r *http.Request) {

	// read body from request
	var body models.ReqCheckCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// validate login values
	if valid := libs.ValidatePhoneOrEmail(body.LoginValue); !valid {
		utils.HandleBadRequestResponse(w, "invalid login value")
		return
	}

	// get code associated with given login value
	i, err := ch.redisDB.Get(body.LoginValue)
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
