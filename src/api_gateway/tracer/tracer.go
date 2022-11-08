package tracer

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle) opentracing.Tracer {

	tracer := opentracing.GlobalTracer()
	cfg := &config.Configuration{
		ServiceName: "api_gateway",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	//1
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	lc.Append(fx.Hook{

		OnStop: func(ctx context.Context) error {
			return closer.Close()
		},
	})

	return tracer
}
