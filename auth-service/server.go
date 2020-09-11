package main

import (
	"log"

	dlv "microservice-app/auth-service/delivery/grpc"
	nosqlRpo "microservice-app/auth-service/repository/nosql"
	sqlRpo "microservice-app/auth-service/repository/sql"
	ucs "microservice-app/auth-service/usecase"

	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	db, err := sqlRpo.ConnectDB("root", "root", "db_user", "3306", "user", "mysql")
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	client, err := nosqlRpo.ConnectDB("db_auth", "6379", "root")
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer client.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	sr := sqlRpo.NewSQLRepository(db)
	nr := nosqlRpo.NewNoSQLRepository(client)
	u := ucs.NewUsecase(sr, nr)

	// grpc server
	grpcServer := grpc.NewServer()
	dlv.NewDeliveryGrpc(grpcServer, u)
	log.Println("auth service runing on port 9000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	}
}
