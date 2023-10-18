package helpers

import (
	"fmt"
	"golang/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("myKey")

type MyCustomClaims struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCustomClaims{
		user.User_id,
		user.Username,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidasiToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claims, nil
}
