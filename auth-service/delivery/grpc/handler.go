package grpc

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"microservice-app/auth-service/delivery/grpc/proto"
	ucs "microservice-app/auth-service/usecase"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	usecase ucs.Usecase
}

// NewDeliveryGrpc ...
func NewDeliveryGrpc(s *grpc.Server, u ucs.Usecase) {
	authServer := &server{
		usecase: u,
	}
	proto.RegisterAuthServer(s, authServer)
}

// Authentication ...
func (u *server) Authentication(ctx context.Context, void *empty.Empty) (*proto.Credential, error) {
	res := new(proto.Credential)

	base64BasicAuth, err := getAuthorizationHeader(ctx)
	if err != nil {
		log.Printf("Error code G-Aen_g <- %v", err)
		return nil, errors.New("G-Aen_g")
	}
	*base64BasicAuth = strings.Replace(*base64BasicAuth, "Basic ", "", -1)
	decodedBasicAuth, err := base64.StdEncoding.DecodeString(*base64BasicAuth)
	stringBasicAuth := string(decodedBasicAuth)
	i := strings.Index(stringBasicAuth, ":")
	username, password := stringBasicAuth[0:i], stringBasicAuth[i+1:]

	t, err := u.usecase.Authentication(username, password)
	if err != nil {
		log.Printf("Error code G-Aen_A <- %v", err)
		return nil, errors.New("G-Aen_A")
	}

	res.Jwt = *t

	log.Printf("### Succesfully Authentication ###")
	return res, nil
}

// Authorization ...
func (u *server) Authorization(ctx context.Context, void *empty.Empty) (*proto.Identity, error) {
	i := new(proto.Identity)

	unsignedToken, err := getAuthorizationHeader(ctx)
	if err != nil {
		log.Printf("Error code G-Aor_g <- %v", err)
		return nil, errors.New("G-Aor_g")
	}
	*unsignedToken = strings.Replace(*unsignedToken, "Bearer ", "", -1)

	c, err := u.usecase.Authorization(*unsignedToken)
	if err != nil {
		log.Printf("Error code G-Aor_A <- %v", err)
		return nil, errors.New("G-Aor_A")
	}

	i.Name = c.Name
	i.Email = c.Email
	i.Address = c.Address
	i.Role = c.Role

	log.Printf("### Succesfully Authorization ###")
	return i, nil
}

func getAuthorizationHeader(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error code G-g : %v", ok)
		return nil, errors.New("G-g")
	}
	arrayOfMd := md.Get("authorization")
	authorization := arrayOfMd[0]

	return &authorization, nil
}
