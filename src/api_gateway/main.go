package main

import (
	"context"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"monorepo/src/api_gateway/routers"
	"monorepo/src/libs/log"
	"monorepo/src/libs/tracer"
	"net/http"

	"github.com/opentracing/opentracing-go"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewMux(lc fx.Lifecycle, conf *configs.Configuration) *tracing.TracedServeMux {
	metricsFactory := jexpvar.NewFactory(10) // 10 buckets for histograms
	logger, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)

	zapLogger := logger.With(zap.String("service", "api_gateway"))
	tracer := tracer.Init("api_gateway", metricsFactory, log.NewFactory(zapLogger))
	opentracing.SetGlobalTracer(tracer)
	handlerWithTracer := tracing.NewServeMux(tracer, conf)

	server := &http.Server{
		Addr:    conf.HTTPPort,
		Handler: handlerWithTracer,
	}
	lc.Append(fx.Hook{

		OnStart: func(context.Context) error {
			// log.Logger.Info("Starting HTTP server.")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// logger.Info("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return handlerWithTracer
}

func main() {

	app := fx.New(
		fx.Provide(
			configs.Config,
			NewMux,
			handlers.New,
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
