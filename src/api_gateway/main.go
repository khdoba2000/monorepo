package main

import (
	"context"
	"log"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/logger"
	"monorepo/src/api_gateway/middleware"
	"monorepo/src/api_gateway/routers"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewMux(lc fx.Lifecycle, logger *log.Logger) *mux.Router {
	logger.Print("Executing NewMux.")
	// Config
	conf := configs.Config()
	if err := conf.Validate(); err != nil {
		panic(err)
	}

	// First, we construct the mux and server. We don't want to start the server
	// until all handlers are registered.
	root := mux.NewRouter()
	root.Use(middleware.Tracer)
	root.Use(middleware.PanicRecovery)
	root.Use(middleware.Logging)
	// casbinJWTRoleAuthorizer, err := middleware.NewCasbinJWTRoleAuthorizer(conf)
	// if err != nil {
	// 	logger.Fatal("Could not initialize Cabin JWT Role Authorizer", zap.Error(err))
	// }
	// root.Use(casbinJWTRoleAuthorizer.Middleware)

	server := &http.Server{
		Addr:    ":8083",
		Handler: root,
	}
	lc.Append(fx.Hook{

		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return root
}

func main() {

	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			handlers.New,
			NewMux,
		),

		fx.Invoke(
			routers.RegisterAuthRoutes,
			routers.RegisterCustomerRoutes,
		),

		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	app.Run()
}
