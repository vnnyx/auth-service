package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vnnyx/auth-service/internal/infrastructure"
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

func (uc *AuthUCImpl) Logout(ctx context.Context, token string) (bool, error) {
	decodedToken, err := uc.decodeToken(token)
	if err != nil {
		return false, err
	}

	err = uc.authRepository.DeleteToken(ctx, decodedToken.AccessUUID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (uc *AuthUCImpl) HasAccess(ctx context.Context, token string) (bool, error) {
	decodedToken, err := uc.decodeToken(token)
	if err != nil {
		return false, err
	}

	got, err := uc.authRepository.GetToken(ctx, decodedToken.AccessUUID)
	if err != nil || got == "" {
		return false, err
	}
	return true, nil
}

func (uc *AuthUCImpl) decodeToken(tokenString string) (decodedToken *model.DecodedStructure, err error) {
	cfg := infrastructure.NewConfig()
	parsePublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cfg.PublicKey))
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return parsePublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	byteData, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteData, &decodedToken); err != nil {
		return nil, err
	}

	return decodedToken, nil
}
