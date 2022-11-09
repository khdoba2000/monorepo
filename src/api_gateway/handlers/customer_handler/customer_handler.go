package customer_handler

import (
	"log"
	"monorepo/src/api_gateway/utils"
	"monorepo/src/libs/etc"
	"monorepo/src/libs/redis"
	libs "monorepo/src/libs/utils"
	"net/http"
	"time"

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

	// Inject the client span context into the headers
	//3
	ch.tracer.Inject(sendCodeHandlerSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	time.Sleep(2 * time.Second)

	// ctx:=opentracing.ContextWithSpan(r.Context(), testHandlerSpan)
	// send this ctx to services called here

	// read body from request
	var body utils.ReqSendCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// generate one-time-password
	code := etc.GenerateCode(4, true)

	// check if request was sent by phone number
	if body.Phone != "" {
		// send code to the number
		if err := libs.SendCodeToPhone(body.Phone, code); err != nil {
			utils.HandleBadRequestErrWithMessage(w, err, "invalid phone number")
			return
		}
		// set number and code to the redis
		if err := ch.redisDB.Set(body.Phone, code); err != nil {
			utils.HandleInternalWithMessage(w, err, "error in setting data to redis")
			return
		}
		// check if request was sent by email
	} else if body.Email != "" {
		// send code to the numbe
		if err := libs.SendCodeToEmail(body.Email, code); err != nil {
			utils.HandleBadRequestErrWithMessage(w, err, "invalid email")
			return
		}
		// set number and code to the redis
		if err := ch.redisDB.Set(body.Email, code); err != nil {
			utils.HandleInternalWithMessage(w, err, "error in setting data to redis")
			return
		}
		// if both body and email are empty thow an error
	} else if body.Email == "" && body.Phone == "" {
		utils.HandleBadRequestResponse(w, "neither email nor phone number is provided")
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

	// Inject the client span context into the headers
	//3
	ch.tracer.Inject(verfyCodeHandlerSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	time.Sleep(2 * time.Second)

	// ctx:=opentracing.ContextWithSpan(r.Context(), testHandlerSpan)
	// send this ctx to services called here

	// read body from request
	var body utils.ReqCheckCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	// check if request was sent by phone number
	if body.Phone != "" {
		// get code associated with given phone number
		i, err := ch.redisDB.Get(body.Phone)
		if err != nil {
			utils.HandleInternalWithMessage(w, err, "error in reading from redis")
			return
		}

		// if the phone number is incorrect or wasn't set before throw error
		v, ok := i.(string)
		if !ok || v == "" {
			utils.HandleBadRequestResponse(w, "phone number is incorrect")
			return
		}
		// if the code is incorrect throw error
		if body.Code != v || body.Code != "7777" {
			utils.HandleBadRequestResponse(w, "code is incorrect")
			return
		}
		// check if request was sent by phone number
	} else if body.Email != "" {
		// get code associated with given email
		i, err := ch.redisDB.Get(body.Email)
		if err != nil {
			utils.HandleInternalWithMessage(w, err, "error in reading from redis")
			return
		}
		// if the email is incorrect or wasn't set before throw error
		v, ok := i.(string)
		if !ok || v == "" {
			utils.HandleBadRequestResponse(w, "email is incorrect")
			return
		}
		// if the code is incorrect throw error
		if body.Code != v || body.Code != "7777" {
			utils.HandleBadRequestResponse(w, "code is incorrect")
			return
		}
		// if both body and email are empty thow an error
	} else if body.Email == "" && body.Phone == "" {
		utils.HandleBadRequestResponse(w, "neither email nor phone number is provided")
		return
	}

	w.WriteHeader(http.StatusOK)

}
