package handlers

import (
	"log"
	"monorepo/src/api_gateway/handlers/auth_handler"
	"monorepo/src/api_gateway/handlers/customer_handler"

	"github.com/opentracing/opentracing-go"
)

// Handlers ...
type Handlers struct {
	AuthHandlers     auth_handler.AuthHandlers
	CustomerHandlers customer_handler.CustomerHandlers
}

// New creates handler
func New(logger *log.Logger, tracer opentracing.Tracer) (*Handlers, error) {

	return &Handlers{
		AuthHandlers:     auth_handler.New(logger, tracer),
		CustomerHandlers: customer_handler.New(logger, tracer),
	}, nil
}
