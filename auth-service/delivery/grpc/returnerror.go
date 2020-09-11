package grpc

import (
	"fmt"
	"microservice-app/auth-service/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func rpcError(codes codes.Code, notifcode, errcode string, err error) error {
	return status.Error(codes, fmt.Sprintf("%v", model.Log(notifcode, errcode, err)))
}
