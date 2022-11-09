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
func New(logger *log.Logger, tracer opentracing.Tracer, rds redis.InMemoryStorageI) CustomerHandlers {
	return &customerHandler{
		tracer:  tracer,
		redisDB: rds,
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

	var body utils.ReqSendCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	code := etc.GenerateCode(4, true)

	if body.Phone != "" {
		if err := libs.SendCodeToPhone(body.Phone, code); err != nil {
			utils.HandleBadRequestErrWithMessage(w, err, "invalid phone number")
			return
		}
		if err := ch.redisDB.Set(body.Phone, code); err != nil {
			utils.HandleInternalWithMessage(w, err, "error in setting data to redis")
			return
		}
	} else if body.Email != "" {
		if err := libs.SendCodeToEmail(body.Email, code); err != nil {
			utils.HandleBadRequestErrWithMessage(w, err, "invalid email")
			return
		}
		if err := ch.redisDB.Set(body.Email, code); err != nil {
			utils.HandleInternalWithMessage(w, err, "error in setting data to redis")
			return
		}
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

	var body utils.ReqCheckCode
	if err := utils.BodyParser(r, body); err != nil {
		utils.HandleBadRequestErrWithMessage(w, err, "invalid body")
		return
	}

	if body.Phone != "" {
		i, err := ch.redisDB.Get(body.Phone)
		if err != nil {
			utils.HandleInternalWithMessage(w, err, "error in reading from redis")
			return
		}
		v, ok := i.(string)
		if !ok || v == "" {
			utils.HandleBadRequestResponse(w, "phone number is incorrect")
			return
		}
		if body.Code != v || body.Code != "7777" {
			utils.HandleBadRequestResponse(w, "code is incorrect")
			return
		}

	} else if body.Email != "" {
		i, err := ch.redisDB.Get(body.Email)
		if err != nil {
			utils.HandleInternalWithMessage(w, err, "error in reading from redis")
			return
		}
		v, ok := i.(string)
		if !ok || v == "" {
			utils.HandleBadRequestResponse(w, "email is incorrect")
			return
		}
		if body.Code != v || body.Code != "7777" {
			utils.HandleBadRequestResponse(w, "code is incorrect")
			return
		}
	} else if body.Email == "" && body.Phone == "" {
		utils.HandleBadRequestResponse(w, "neither email nor phone number is provided")
		return
	}

	w.WriteHeader(http.StatusOK)

}
