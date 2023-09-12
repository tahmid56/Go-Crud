package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context){
	if len(c.Request.Header["Authorization"]) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User Not Authorized",
		})
		return
	}

	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You should provide your own secret key here to verify the token
		// Replace `YourSecretKey` with your actual secret key
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Set("userid", claims["sub"])
	fmt.Println(claims["sub"])
	
	c.Next()
}