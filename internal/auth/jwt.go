package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

func GenerateToken(jwtSecret string, user model.User) (string, error) {
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
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
