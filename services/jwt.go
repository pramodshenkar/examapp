package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(1440 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
func DecodeToken(tkStr string) (string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}
	if !tkn.Valid {
		return "", err
	}
	return claims.UserID, nil
}
