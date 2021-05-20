package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/plyama/auth/internal/models"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ReadToken(r *http.Request) (*jwt.Token, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return &jwt.Token{}, errors.New("there is no auth header")
	}

	headerWords := strings.Split(authHeader, " ")
	if len(headerWords) != 2 {
		return &jwt.Token{}, errors.New("not valid auth header")
	}

	authToken := headerWords[1]

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("not valid token")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	return token, err
}
