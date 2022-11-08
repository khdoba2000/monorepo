package customer_handler

import (
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
)

type CustomerHandlers interface {
	TestHandler2(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	tracer opentracing.Tracer
}

// New creates handler
func New(logger *log.Logger, tracer opentracing.Tracer) CustomerHandlers {
	return &customerHandler{tracer: tracer}
}
func (ch *customerHandler) TestHandler2(w http.ResponseWriter, r *http.Request) {
	// ch.logger.Print("Got a request.")
	w.Write([]byte("Hello, World2!"))
}
