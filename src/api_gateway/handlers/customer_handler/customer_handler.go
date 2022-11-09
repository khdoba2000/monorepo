package customer_handler

import (
	"log"
	"monorepo/src/api_gateway/models"
	"monorepo/src/api_gateway/utils"
	"monorepo/src/libs/etc"
	"monorepo/src/libs/redis"
	libs "monorepo/src/libs/utils"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type CustomerHandlers interface {
	TestHandler2(w http.ResponseWriter, r *http.Request)
	VerfyCodeHandler(w http.ResponseWriter, r *http.Request)
	SendCodeHandler(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	tracer  opentracing.Tracer
	redisDB redis.InMemoryStorageI
}

// New creates handler
func New(logger *log.Logger, tracer opentracing.Tracer) CustomerHandlers {
	return &customerHandler{
		tracer: tracer,
		// redisDB: rds,
	}
}
func (ch *customerHandler) TestHandler2(w http.ResponseWriter, r *http.Request) {
	// ch.logger.Print("Got a request.")
	w.Write([]byte("Hello, World2!"))
}

func (ch *customerHandler) SendCodeHandler(w http.ResponseWriter, r *http.Request) {
	sendCodeHandlerSpan := ch.tracer.StartSpan("SendCodeHandler")
	defer sendCodeHandlerSpan.Finish()

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
	ext.SpanKindRPCClient.Set(sendCodeHandlerSpan)
	ext.HTTPUrl.Set(sendCodeHandlerSpan, r.URL.Path)
	ext.HTTPMethod.Set(sendCodeHandlerSpan, "GET")

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
	verfyCodeHandlerSpan := ch.tracer.StartSpan("VerfyCodeHandler")
	defer verfyCodeHandlerSpan.Finish()

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
	ext.SpanKindRPCClient.Set(verfyCodeHandlerSpan)
	ext.HTTPUrl.Set(verfyCodeHandlerSpan, r.URL.Path)
	ext.HTTPMethod.Set(verfyCodeHandlerSpan, "GET")

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
