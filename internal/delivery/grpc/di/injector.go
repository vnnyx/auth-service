//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/vnnyx/auth-service/internal/delivery/grpc"
	"github.com/vnnyx/auth-service/internal/infrastructure"
	"github.com/vnnyx/auth-service/internal/repository"
	"github.com/vnnyx/auth-service/internal/usecase"
)

func InitializeGRPCAuthServer() *grpc.Server {
	wire.Build(
		infrastructure.NewRedisClient,
		repository.NewAuthRepository,
		usecase.NewAuthUC,
		grpc.NewAuthServiceServer,
	)
	return nil
}
