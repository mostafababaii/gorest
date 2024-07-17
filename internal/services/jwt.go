package services

import (
	"errors"
	"fmt"
	"github.com/mostafababaii/gorest/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mostafababaii/gorest/config"
)

type JwtToken struct {
	Secret   []byte
	LifeSpan time.Duration
}

func NewJwtToken() *JwtToken {
	return &JwtToken{
		Secret:   config.JwtSecret,
		LifeSpan: config.JwtTokenLifeSpan,
	}
}

func (t *JwtToken) Generate(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userID"] = user.ID
	claims["exp"] = time.Now().Add(t.LifeSpan).Unix()

	tokenString, err := token.SignedString(t.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *JwtToken) Validate(token string) (*models.User, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	user := models.User{}
	user.ID = uint(claims["userID"].(float64))
	return &user, nil
}
