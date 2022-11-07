package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func TestSrvice(t *testing.T) {
	go main()

	time.Sleep(time.Second) // Leave time for service to stat
	tracer := opentracing.GlobalTracer()
	cfg := &config.Configuration{
		ServiceName: "client",
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
	time.Sleep(time.Second)

	url := "http://localhost:8080/one/hello"
	req, _ := http.NewRequest("GET", url, nil)

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
	ext.SpanKindRPCClient.Set(clientSpan)
	ext.HTTPUrl.Set(clientSpan, url)
	ext.HTTPMethod.Set(clientSpan, "GET")

	// Inject the client span context into the headers
	//3
	tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := "Hello, World!"
	actual, _ := ioutil.ReadAll(resp.Body)
	if expected != string(actual) {
		fmt.Println("expected", expected)
		fmt.Println("actual:", string(actual))
		t.Fail()
	}

}
