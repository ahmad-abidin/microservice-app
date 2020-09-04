package main

import (
	"context"
	"log"

	proto "microservice-app/auth-service/delivery/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("auth_service:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer conn.Close()

	a := proto.NewAuthClient(conn)

	log.Println("### Authentication ###")
	Identity := proto.Credential{
		Username: "abidin",
		Password: "password123",
	}
	res, err := a.Authentication(context.Background(), &Identity)
	if err != nil {
		log.Fatalf("error when calling Authentication: %v", err)
	}
	log.Printf("credential: %v", res)
	log.Println("### Successfully Authentication ###")

	log.Println("### Authorization ###")
	credential := proto.Token{
		Jwt: "token",
	}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", res.Jwt)
	res2, err := a.Authorization(ctx, &credential)
	if err != nil {
		log.Fatalf("error when calling Authorization: %v", err)
	}
	log.Printf("full_identity: %v", res2)
	log.Println("### Successfully Authorization ###")
}
