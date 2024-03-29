package repository

import (
	"context"

	"github.com/vnnyx/auth-service/internal/model"
)

type AuthRepository interface {
	StoreToken(ctx context.Context, details model.TokenDetails) error
	DeleteToken(ctx context.Context, accessUuid string) error
	GetToken(ctx context.Context, accessUuid string) (access string, err error)
}
