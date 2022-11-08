package main

import (
	"net"

	"google.golang.org/grpc/reflection"
	"monorepo/src/auth_service/configs"
	"monorepo/src/auth_service/pkg/db"
	"monorepo/src/auth_service/service"
	"monorepo/src/auth_service/storage"
	pb "monorepo/src/idl/auth_service"
	"monorepo/src/libs/logger"

	"google.golang.org/grpc"
)

func main() {

	//Loaf configurations
	config := configs.Config()

	//Create logger instance
	log := logger.New(config.LogLevel, "AuthService")

	//Cleanup logger when returning main func
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	//Initialize database, make a connection with postgres
	connDB, err := db.Init(*config)

	if err != nil {
		log.Fatal("grpc failed to listen: ", logger.Error(err))
	}

	log.Info("authService: sqlxConfig",
		logger.String("host", config.PostgresHost),
		logger.Int("port", config.PostgresPort),
		logger.String("database", config.PostgresDatabase),
	)

	// Create storage instance
	storage := storage.New(connDB)

	// Make an authentication service instance
	authService := service.New(storage, *config, log)

	// Create grpc server, registering and listening
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":8084")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", logger.Error(err))
	}

	log.Info("auth server running on port : ", logger.String("port", "8084"))

}
