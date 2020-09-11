package grpc

import (
	"context"
	"errors"
	"fmt"
	"microservice-app/auth-service/delivery/grpc/proto"
	"microservice-app/auth-service/model"
	ucs "microservice-app/auth-service/usecase"

	"google.golang.org/grpc/codes"

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
		return nil, rpcError(codes.FailedPrecondition, "e", "grpc-Aen_gAH", err)
	}

	t, err := u.usecase.Authentication(*base64BasicAuth)
	if err != nil {
		return nil, rpcError(codes.PermissionDenied, "e", "grpc-Aen_Aen", err)
	}

	res.Jwt = *t

	model.Log("s", "grpc-Aor", errors.New("Successfully Authentication"))

	return res, nil
}

// Authorization ...
func (u *server) Authorization(ctx context.Context, void *empty.Empty) (*proto.Identity, error) {
	i := new(proto.Identity)

	unsignedToken, err := getAuthorizationHeader(ctx)
	if err != nil {
		return nil, rpcError(codes.PermissionDenied, "e", "grpc-Aor_gAH", err)
	}

	c, err := u.usecase.Authorization(*unsignedToken)
	if err != nil {
		return nil, rpcError(codes.PermissionDenied, "e", "grpc-Aor_Aor", err)
	}

	i.Name = c.Name
	i.Email = c.Email
	i.Address = c.Address
	i.Role = c.Role

	model.Log("s", "grpc-Aor", errors.New("Successfully Athorization"))

	return i, nil
}

func getAuthorizationHeader(ctx context.Context) (*string, error) {
	var authorization string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, rpcError(codes.PermissionDenied, "e", "grpc-gAH", fmt.Errorf("authorization header : %v", ok))
	}

	arrayOfMd := md.Get("authorization")
	if len(arrayOfMd) == 0 {
		return nil, rpcError(codes.PermissionDenied, "e", "grpc-gAH_G", errors.New("authorization header is nil"))
	}

	authorization = arrayOfMd[0]

	return &authorization, nil
}
