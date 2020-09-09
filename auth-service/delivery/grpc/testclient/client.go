package main

import (
	"context"
	"log"

	proto "microservice-app/auth-service/delivery/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial("auth_service:9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect : %v", err)
	}
	defer conn.Close()
	a := proto.NewAuthClient(conn)

	log.Println("### Authentication ###")

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Basic YWhtYWQuYWJpZGluQG1haWwuY29tOnBhc3N3b3JkMTIzNA==")
	token, err := a.Authentication(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("error when calling Authentication: %v", err)
	}
	log.Printf("credential: %v", token)
	log.Println("### Successfully Authentication ###")

	log.Println("### Authorization ###")
	ctx = metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token.Jwt)
	identity, err := a.Authorization(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("error when calling Authorization: %v", err)
	}
	log.Printf("full_identity: %v", identity)
	log.Println("### Successfully Authorization ###")
}
