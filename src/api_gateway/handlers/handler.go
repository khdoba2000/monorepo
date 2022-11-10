package handlers

import (
	"monorepo/src/api_gateway/handlers/auth_handler"
	"monorepo/src/api_gateway/handlers/customer_handler"
	"monorepo/src/libs/log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Handlers ...
type Handlers struct {
	AuthHandlers     auth_handler.AuthHandlers
	CustomerHandlers customer_handler.CustomerHandlers
}

// New creates handler
func New() (*Handlers, error) {
	zapLogger, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	logger := log.NewFactory(zapLogger)
	return &Handlers{
		AuthHandlers:     auth_handler.New(logger),
		CustomerHandlers: customer_handler.New(logger),
	}, nil
}
