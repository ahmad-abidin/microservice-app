package main

import (
	"log"
	"microservice-app/auth-service/proto"
	"microservice-app/auth-service/service"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := service.Server{}
	proto.RegisterAuthServer(grpcServer, &s)
	log.Println("auth service runing on port 9000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	}
}
