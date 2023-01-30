package commen

import (
	"awesomeProject1/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var MySigningKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User) (string, error) {
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})
	return token, claims, err
}
