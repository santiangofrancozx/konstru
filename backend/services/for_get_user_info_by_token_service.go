package services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetUserInfoByTokenService() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			fmt.Println("Failed to get token from cookie:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := Adapters.DecodeJWTToken(tokenString)
		if err != nil {
			fmt.Println("Failed to decode JWT token:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		exp, ok := claims["exp"].(float64)
		if !ok || float64(time.Now().Unix()) > exp {
			fmt.Println("Token has expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			fmt.Println("Failed to extract email from token claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := Adapters.SelectUserByUsernameInfo(email)
		if err != nil {
			fmt.Println("Failed to fetch user info from database:", err)
			c.JSON(http.StatusNotFound, nil)
			return
		}
		// Respond with user info
		c.JSON(http.StatusOK, user)

	}

}
