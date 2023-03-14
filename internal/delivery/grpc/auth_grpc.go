package grpc

import (
	"context"

	"github.com/vnnyx/auth-service/internal/model"
	"github.com/vnnyx/auth-service/internal/usecase"
	pb "github.com/vnnyx/auth-service/pb/auth"
)

type Server struct {
	authUC usecase.AuthUC
	pb.UnimplementedAuthServiceServer
}

func NewGRPCServer(authUC usecase.AuthUC) *Server {
	return &Server{authUC: authUC}
}

func (s *Server) Login(ctx context.Context, req *pb.AuthRequest) (*pb.Token, error) {
	res, err := s.authUC.Login(ctx, model.User{ID: req.User.Id, Username: req.User.Username}, model.LoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return res.ToGRPCResponse(), nil
}
