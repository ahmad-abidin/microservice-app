package grpc

import (
	"context"
	"errors"
	"log"
	"microservice-app/auth-service/delivery/grpc/proto"
	"microservice-app/auth-service/model"
	"microservice-app/auth-service/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	usecase usecase.Usecase
}

// NewDeliveryGrpc ...
func NewDeliveryGrpc(s *grpc.Server, u usecase.Usecase) {
	authServer := &server{
		usecase: u,
	}
	proto.RegisterAuthServer(s, authServer)
}

// Authentication ...
func (u *server) Authentication(ctx context.Context, c *proto.Credential) (res *proto.Token, err error) {
	m := model.Credential{}
	// res := new(proto.Token)

	m.Username = c.Username
	m.Password = c.Password

	t, err := u.usecase.Authentication(m)
	if err != nil {
		log.Fatalf("Error code G-AenA <- %v", err)
		return nil, errors.New("G-AenA")
	}

	res.Jwt = *t

	log.Printf("### Succesfully Authentication ###")
	return res, nil
}

// Authorization ...
func (u *server) Authorization(ctx context.Context, t *proto.Token) (*proto.Identity, error) {
	i := new(proto.Identity)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalf("Error code G-AF <- %v", ok)
		return nil, errors.New("G-AF")
	}
	arrayOfMd := md.Get("authorization")
	unsignedToken := arrayOfMd[0]

	c, err := u.usecase.Authorization(unsignedToken)
	if err != nil {
		log.Fatalf("Error code G-AorA <- %v", err)
		return nil, errors.New("G-AorA")
	}

	i.Name = c.Name
	i.Email = c.Email
	i.Address = c.Address

	log.Printf("### Succesfully Authorization ###")
	return i, nil
}
