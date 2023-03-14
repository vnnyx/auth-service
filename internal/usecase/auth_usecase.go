package usecase

import (
	"context"

	"github.com/vnnyx/auth-service/internal/model"
)

type AuthUC interface {
	Login(ctx context.Context, user model.User, request model.LoginRequest) (response model.LoginResponse, err error)
}
