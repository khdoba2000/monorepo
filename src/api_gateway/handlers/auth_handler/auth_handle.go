package auth_handler

import (
	"log"
	"net/http"
	//"monorepo/src/api_gateway/ci"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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

func (ah *authHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	// container := ci.Get()
}
