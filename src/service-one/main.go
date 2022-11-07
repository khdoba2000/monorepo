package main

import (
	"fmt"
	"net/http"

	"monorepo/src/libs/hello"

	"github.com/gin-gonic/gin"

	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func main() {

	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "your_service_name",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	r := gin.Default()
	r.GET("/one/hello", func(c *gin.Context) {

		//2
		clientSpan := tracer.StartSpan("clientspan")
		defer clientSpan.Finish()
		// Extract the context from the headers
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		serverSpan := tracer.StartSpan("server", ext.RPCServerOption(spanCtx))
		defer serverSpan.Finish()
		time.Sleep(time.Second)
		c.String(http.StatusOK, hello.Greet("World"))
	})
	_ = r.Run(":8080")
	fmt.Println("request recieved to service one")
}
