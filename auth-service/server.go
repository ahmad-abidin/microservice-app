package main

import (
	"log"

	d "microservice-app/auth-service/delivery/grpc"
	"microservice-app/auth-service/repository"
	"microservice-app/auth-service/usecase"

	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	db, err := repository.ConnectDB("root", "root", "db_user", "3306", "user", "mysql")
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	r := repository.NewRepository(db)
	s := usecase.NewUsecase(r)

	// grpc
	grpcServer := grpc.NewServer()
	d.NewDeliveryGrpc(grpcServer, s)
	log.Println("auth service runing on port 9000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	}
}
