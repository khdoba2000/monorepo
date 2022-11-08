package middleware

import (
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// Tracer logs the time
func Tracer(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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
		defer closer.Close()
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		}

		//2
		clientSpan := tracer.StartSpan("clientspan")
		defer clientSpan.Finish()

		// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
		ext.SpanKindRPCClient.Set(clientSpan)
		ext.HTTPUrl.Set(clientSpan, req.URL.Path)
		ext.HTTPMethod.Set(clientSpan, "GET")

		// Inject the client span context into the headers
		//3
		tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

		next.ServeHTTP(w, req)
		return
	})
}
