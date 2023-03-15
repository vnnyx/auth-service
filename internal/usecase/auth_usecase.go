package usecase

import (
	"context"

	"github.com/vnnyx/auth-service/internal/model"
)

type AuthUC interface {
	Login(ctx context.Context, user model.User, request model.LoginRequest) (response model.LoginResponse, err error)
	HasAccess(ctx context.Context, token string) (bool, error)
	Logout(ctx context.Context, token string) (bool, error)
}
