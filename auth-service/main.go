package main

import (
	"log"

	dlv "microservice-app/auth-service/delivery/api"
	"microservice-app/auth-service/model"
	nosqlRpo "microservice-app/auth-service/repository/nosql"
	sqlRpo "microservice-app/auth-service/repository/sql"
	ucs "microservice-app/auth-service/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
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

	sr := sqlRpo.NewSQLRepository(db)
	nr := nosqlRpo.NewNoSQLRepository(client)
	u := ucs.NewUsecase(sr, nr)

	// api
	apiServer := echo.New()
	dlv.NewDeliveryAPI(apiServer, u)
	model.Log("e", "error when start api server", apiServer.Start(":9000"))

	// grpc server
	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen on port 9000: %v", err)
	// }
	// grpcServer := grpc.NewServer()
	// dlv.NewDeliveryGrpc(grpcServer, u)
	// log.Println("auth service runing on port 9000")
	// err = grpcServer.Serve(lis)
	// if err != nil {
	// 	log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	// }
}
