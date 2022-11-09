package auth_handler

import (
	"encoding/json"
	"log"
	"monorepo/src/api_gateway/ci"
	"monorepo/src/api_gateway/handlers/models"
	"monorepo/src/api_gateway/utils"
	"monorepo/src/libs/etc"
	"net/http"
	"strings"

	//"monorepo/src/api_gateway/ci"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
	SendCode(w http.ResponseWriter, r *http.Request)
	Verify(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	tracer opentracing.Tracer
}

// New creates auth handlers
func New(logger *log.Logger, tracer opentracing.Tracer) AuthHandlers {
	return &authHandler{tracer: tracer}
}

func (ah *authHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	testHandlerSpan := ah.tracer.StartSpan("TestHandler")
	defer testHandlerSpan.Finish()

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
	ext.SpanKindRPCClient.Set(testHandlerSpan)
	ext.HTTPUrl.Set(testHandlerSpan, r.URL.Path)
	ext.HTTPMethod.Set(testHandlerSpan, "GET")

	// Inject the client span context into the headers
	//3
	ah.tracer.Inject(testHandlerSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	time.Sleep(2 * time.Second)

	// ctx:=opentracing.ContextWithSpan(r.Context(), testHandlerSpan)
	// send this ctx to services called here

	w.Write([]byte("Hello, World1!"))
}

func (ah *authHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	container := ci.Get()

	var (
		body models.SendSMSModel
	)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !utils.InEnums(body.Type, []string{"signup", "forgot", "reset"}) {
		writeError(w, "Invalid type", http.StatusBadRequest)
		return
	}

	if !strings.Contains(body.PhoneNumber, "+") || len(body.PhoneNumber) != 13 {
		writeError(w, "Invalid input a phone number. Must be + and number count 13", http.StatusBadRequest)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(configs.Config().CtxTimeout))
	// defer cancel()

	// need to check authService's check field method

	code := etc.GenerateCode(4)
	data := models.RedisData{
		Value:    body.PhoneNumber,
		Code:     code,
		Verified: false,
	}

	bodyJSON, err := json.Marshal(data)

	if err != nil {
		return
	}

	// this code and phone number will be saved only 3 minutes in redis
	err = container.Redis.SetWithTTL(body.PhoneNumber, string(bodyJSON), 180)
	if err != nil {
		writeError(w, "Error at write to redis. in SetWithTTL", http.StatusInternalServerError)
		return
	}

	writeJSON(w, models.SuccessMessage{
		Success: true,
	})
}

func (ah *authHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var (
		body models.VerifyModel
	)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Error parsing body: ", err)
		return
	}

	//Geting code from redis
	user := models.RedisData{}
	userJSON, err := redis.String(ci.Get().Redis.Get(body.PhoneNumber))
	if err != nil {
		log.Println("Error while getting from redis", err)
		writeError(w, "Error while getting from redis", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		writeError(w, "Error while unmarshalling! ", http.StatusBadRequest)
		return
	}

	//Checking whether received code is valid
	if body.Code != "7777" {
		writeError(w, "Code is invalid! ", http.StatusBadRequest)
		return
	}

	err = ci.Get().Redis.SetWithTTL(body.PhoneNumber, "true", 1800)
	if err != nil {
		log.Println("Error putting phone_number and data in redis! ", err)
		writeError(w, "Error while writing to Redis! ", http.StatusBadRequest)
		return
	}

	writeJSON(w, models.SuccessMessage{
		Success: true,
	})
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		body models.LoginModel
	)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !strings.Contains(body.PhoneNumber, "+") || len(body.PhoneNumber) == 13 {
		writeError(w, "phone number is not correctly filled", http.StatusBadRequest)
		return
	}

	err = body.Validate()
	if err != nil {
		writeError(w, "one of the fields is not correct", http.StatusBadRequest)
		return
	}

	// ctx, cancel := context.WithTimeout(context.Background(),
	// 	time.Second*time.Duration(configs.Config().CtxTimeout))
	// defer cancel()

	//Integration with authService's login method

}
