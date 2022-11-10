package main

import (
	"fmt"
	"monorepo/src/libs/log"
	"monorepo/src/libs/tracer"
	"net"

	otgrpc "github.com/opentracing-contrib/go-grpc"

	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"

	pb "monorepo/src/idl/auth_service"

	"monorepo/src/auth_service/configs"
	"monorepo/src/auth_service/pkg/db"
	"monorepo/src/auth_service/service"
	"monorepo/src/auth_service/storage"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {

	//Load configurations
	config := configs.Config()

	//Create logger instance
	// log := logger.New()

	//Cleanup logger when returning main func
	// defer func(l logger.Logger) {
	// 	err := logger.Cleanup(l)
	// 	if err != nil {
	// 		logger.Fatal()
	// 	}
	// }(log)

	//Initialize database, make a connection with postgres
	connDB, err := db.Init(*config)
	if err != nil {
		fmt.Println("failed to connect with db: ", err)
	}

	// logger.Info("authService: sqlxConfig",
	// 	logger.String("host", config.PostgresHost),
	// 	logger.Int("port", config.PostgresPort),
	// 	logger.String("database", config.PostgresDatabase),
	// )

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

	// Make an authentication service instance
	authServer := service.New(storage.New(connDB), logger, tracer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	//listenting tcp rpcport
	lis, err := net.Listen("tcp", config.RPCPort)
	if err != nil {
		fmt.Println("listening tcp error: ", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err)
	}

	fmt.Println("crm server running on port : ", config.RPCPort)

}
