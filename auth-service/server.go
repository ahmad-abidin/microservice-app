package main

import (
	"log"

	dlv "microservice-app/auth-service/delivery/grpc"
	rpo "microservice-app/auth-service/repository/sql"
	ucs "microservice-app/auth-service/usecase"

	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	db, err := rpo.ConnectDB("root", "root", "db_user", "3306", "user", "mysql")
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	r := rpo.NewRepository(db)
	s := ucs.NewUsecase(r)

	// grpc
	grpcServer := grpc.NewServer()
	dlv.NewDeliveryGrpc(grpcServer, s)
	log.Println("auth service runing on port 9000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	}
}
