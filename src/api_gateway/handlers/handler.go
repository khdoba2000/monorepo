package handlers

import (
	"fmt"
	"monorepo/src/api_gateway/handlers/auth_handler"
	"monorepo/src/api_gateway/handlers/customer_handler"
	"monorepo/src/libs/logger"
)

// Handlers ...
type Handlers struct {
	AuthHandlers     auth_handler.AuthHandlers
	CustomerHandlers customer_handler.CustomerHandlers
}

// New creates handler
func New(logger logger.Logger) (*Handlers, error) {
	fmt.Println("handler New")

	return &Handlers{
		AuthHandlers:     auth_handler.New(logger),
		CustomerHandlers: customer_handler.New(logger),
	}, nil
}
