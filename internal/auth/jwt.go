package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucasshuan/b2c-ecommerce-api/configs"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
)

type Permission string

const (
	Guest Permission = "guest"
	User  Permission = "user"
	Admin Permission = "admin"
)

type Claims struct {
	UserID     uint       `json:"userId"`
	Permission Permission `json:"permission"`
	jwt.StandardClaims
}

func GetUserPermission(user model.User) Permission {
	if user.IsAdmin {
		return Admin
	}
	return User
}

func GenerateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	permission := GetUserPermission(user)

	claims := &Claims{
		UserID:     user.ID,
		Permission: permission,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.Config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, jwtSecret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return claims, nil
}
