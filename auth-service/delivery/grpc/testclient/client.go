package main

import (
	"context"
	"log"
	"os"

	proto "microservice-app/auth-service/delivery/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	argsRaw := os.Args

	conn, err := grpc.Dial("auth_service:9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect : %v", err)
	}
	defer conn.Close()

	client := proto.NewAuthClient(conn)

	log.Println("### Authentication ###")
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Basic "+argsRaw[1])
	credential, err := client.Authentication(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("credential: %v", credential)
	log.Println("### Successfully Authentication ###\n\n")

	log.Println("### Authorization ###")
	var bearerToken string
	if credential != nil {
		bearerToken = credential.Jwt
	}
	ctx = metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+bearerToken)
	identity, err := client.Authorization(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("identity: %v", identity)
	log.Println("### Successfully Authorization ###")
}
