package auth_handler

import (
	"context"
	"encoding/json"
	"log"
	"monorepo/src/api_gateway/ci"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/handlers/models"
	"monorepo/src/api_gateway/utils"
	"monorepo/src/libs/etc"
	"net/http"
<<<<<<< HEAD
	//"monorepo/src/api_gateway/ci"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
=======
	"strings"
	"time"

	"go.uber.org/atomic"
>>>>>>> 1320807 (--am)
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
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

	var (
		codeSafe atomic.String
	)

	codeSafe.Store(etc.GenerateCode(4))
	verificationCode := codeSafe

	if !strings.Contains(body.PhoneNumber, "+") || len(body.PhoneNumber) != 13 {
		writeError(w, "Invalid input a phone number. Must be + and number count 13", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(configs.Config().CtxTimeout))
	defer cancel()

}
