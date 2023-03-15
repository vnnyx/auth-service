package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vnnyx/auth-service/internal/infrastructure"
	"github.com/vnnyx/auth-service/internal/model"
)

func CreateToken(user model.User) (*model.TokenDetails, error) {
	cfg := infrastructure.NewConfig()

	access_uuid := uuid.NewString()
	expired := time.Now().Add(time.Minute * time.Duration(cfg.JWTMinutes))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expired
	claims["authorized"] = true
	claims["user"] = user.Username
	claims["user_id"] = user.ID
	claims["access_uuid"] = access_uuid
	tokenString, err := token.SignedString([]byte(cfg.PrivateKey))
	if err != nil {
		return nil, err
	}
	return &model.TokenDetails{
		AccessToken: tokenString,
		AccessUUID:  access_uuid,
		AtExpires:   expired.Unix(),
	}, nil
}
