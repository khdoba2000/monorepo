package main

import (
	"context"
	"fmt"
	"monorepo/src/libs/log"
	"monorepo/src/libs/tracer"
	"net"
	"time"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"

	pb "monorepo/src/idl/auth_service"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
	logger log.Factory
	tracer opentracing.Tracer
}

func (s authServer) CustomerSignUp(ctx context.Context, r *pb.CustomerSignUpRequest) (*pb.AuthResponse, error) {

	s.logger.For(ctx).Info("CustomerSignUp req", zap.String("name", r.Name))

	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := s.tracer.StartSpan("Query database", opentracing.ChildOf(span.Context()))
		span.SetTag("param.name", r.Name)
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	//simulate signup reg
	time.Sleep(1 * time.Second)

	return &pb.AuthResponse{AccessToken: "access", RefreshToken: "refresh"}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":8084")

	if err != nil {
		fmt.Println("grpc failed to listen: ")
		panic(err)
	}

	metricsFactory := jexpvar.NewFactory(10) // 10 buckets for histograms
	loggerForTracer, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)

	zapLogger := loggerForTracer.With(zap.String("service", "auth_service"))
	tracer := tracer.Init("auth_service", metricsFactory, log.NewFactory(zapLogger))

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	logger := log.NewFactory(zapLogger)

	pb.RegisterAuthServiceServer(grpcServer, &authServer{
		logger: logger,
		tracer: tracer,
	})

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("Serve")
	}
}
