package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"monorepo/src/idl/auth_service"
)

type authServer struct {
	auth_service.UnimplementedAuthServiceServer
}

func main() {

	listener, err := net.Listen("tcp", ":8084")

	if err != nil {
		log.Fatalln("grpc failed to listen: ", err)
	}

	log.Println("Successfully created TCP listener")

	grpcServer := grpc.NewServer()
	auth_service.RegisterAuthServiceServer(grpcServer, &authServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
