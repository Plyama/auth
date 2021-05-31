package middlewares

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"

	"github.com/plyama/auth/internal/utils/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type contextKey string

var authContextKey = contextKey("auth")

func AuthorizeViaHeader(c *gin.Context) {
	token, err := auth.ReadToken(c.Request)
	if err != nil {
		switch err.(type) {
		case auth.NotValidAuthHeader:
			log.Printf("not valid auth header: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
		default:
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	if !token.Valid {
		log.Println("not valid token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("not ok claims")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	value, ok := claims["user_data"]
	if !ok {
		log.Println("there is no user data")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var userData auth.UserData
	err = mapstructure.Decode(value, &userData)
	if err != nil {
		log.Println("failed to transform to auth.UserData")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(c.Request.Context(), authContextKey, userData)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func GetUserData(ctx context.Context) (*auth.UserData, error) {
	data, ok := ctx.Value(authContextKey).(auth.UserData)
	if !ok {
		return nil, errors.New("failed to get user's ID")
	}

	return &data, nil
}
