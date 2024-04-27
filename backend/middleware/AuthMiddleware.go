package middleware

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if token == nil {
		// Token doesn't exist, respond with a message
		fmt.Println("Token does not exist")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// If the token is about to expire, generate a new token
			newToken, err := generateToken(claims["email"].(string))
			if err != nil {
				fmt.Println("Error generating new token:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Set the new token in the response
			c.Header("Authorization", "Bearer "+newToken)
		}

		// Find the user with token Subject
		var user models.Usuario
		user, err = Adapters.SelectUserByUsername(claims["email"].(string))

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		c.Set("user", user)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func generateToken(email string) (string, error) {
	// Create a new token with updated expiration time
	expirationTime := time.Now().Add(time.Minute * 15) // Set expiration to 15 minute from now
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
