package service

import (
	"context"

	"github.com/ahmad-abidin/microservice-app/auth-service/model"
)

type Server struct{}

func (s *Server) Authentication(ctx context.Context, identity *model.Identity) {

}
