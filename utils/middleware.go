package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		err := tokenValid(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		userId, err := extractTokenID(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("can't extract token: %s", err.Error())})
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func tokenValid(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func extractTokenID(tokenString string) (primitive.ObjectID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return primitive.NilObjectID, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userIdStr, ok := claims["user_id"].(string)
		if !ok {
			return primitive.NilObjectID, fmt.Errorf("user id claim is not a string")
		}
		userId, err := primitive.ObjectIDFromHex(userIdStr)
		if err != nil {
			return primitive.NilObjectID, err
		}
		return userId, nil
	}
	return primitive.NilObjectID, nil
}
