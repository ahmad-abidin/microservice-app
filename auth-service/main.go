package main

import (
	"log"

	dlv "microservice-app/auth-service/delivery/api"
	"microservice-app/auth-service/model"
	nosqlRpo "microservice-app/auth-service/repository/nosql"
	sqlRpo "microservice-app/auth-service/repository/sql"
	ucs "microservice-app/auth-service/usecase"

	_ "github.com/go-sql-driver/mysql"
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

	// api server -> import 	dlv "microservice-app/auth-service/delivery/api"
	apiServer := dlv.NewDeliveryAPI(u)
	model.Log("e", "error when start API Server", apiServer.Start(":9000"))

	// grpc server -> import 	dlv "microservice-app/auth-service/delivery/grpc"
	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen on port 9000: %v", err)
	// }
	// grpcServer := dlv.NewDeliveryGrpc(u)
	// model.Log("e", "error when start API Server", grpcServer.Serve(lis))
}
