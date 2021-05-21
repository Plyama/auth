package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/plyama/auth/internal/utils/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type contextKey string

var authContextKey = contextKey("auth")

func Authorize(c *gin.Context) {
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

	id, ok := claims["id"].(float64)
	if !ok {
		log.Println("there is no id in claims")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(c.Request.Context(), authContextKey, int(id))
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func GetUserID(ctx context.Context) (int, error) {
	ID, ok := ctx.Value(authContextKey).(int)
	if !ok {
		return 0, errors.New("failed to get user's ID")
	}

	return ID, nil
}
