package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vnnyx/auth-service/internal/infrastructure"
	"github.com/vnnyx/auth-service/internal/model"
)

func CreateToken(user model.User) (td *model.TokenDetails, err error) {
	cfg := infrastructure.NewConfig()

	expired := time.Now().Add(10 * time.Minute)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expired
	claims["authorized"] = true
	claims["user"] = user.Username
	claims["id"] = user.ID
	tokenString, err := token.SignedString([]byte(cfg.PrivateKey))
	if err != nil {
		return td, err
	}

	td.AccessToken = tokenString
	td.AccessUUID = uuid.New().String()
	td.AtExpires = expired.Unix()
	return td, nil
}
