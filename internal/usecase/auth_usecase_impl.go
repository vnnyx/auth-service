package usecase

import (
	"context"

	"github.com/vnnyx/auth-service/internal/model"
	"github.com/vnnyx/auth-service/internal/repository"
	"github.com/vnnyx/auth-service/internal/util"
)

type AuthUCImpl struct {
	authRepository repository.AuthRepository
}

func NewAuthUC(authRepository repository.AuthRepository) AuthUC {
	return &AuthUCImpl{authRepository: authRepository}
}

func (uc *AuthUCImpl) Login(ctx context.Context, user model.User, request model.LoginRequest) (response model.LoginResponse, err error) {
	td, err := util.CreateToken(model.User{
		ID:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		return response, err
	}

	tokenDetails := &model.TokenDetails{
		AccessToken: td.AccessToken,
		AccessUUID:  td.AccessUUID,
		AtExpires:   td.AtExpires,
	}

	err = uc.authRepository.StoreToken(ctx, *tokenDetails)
	if err != nil {
		return response, err
	}

	response = model.LoginResponse{
		AccessToken: td.AccessToken,
	}

	return response, nil
}
