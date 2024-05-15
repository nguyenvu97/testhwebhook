package Token

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"strings"
	"time"
)

func GenerateToken(email string) (*string, error) {
	// Tạo một đối tượng Claims cho JWT
	claims := jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	log.Info(time.Now().Add(time.Second).Unix())

	privateKey := os.Getenv("PRIVATE_KEY")
	fmt.Println("Private Key:", privateKey)
	if privateKey == "" {
		return nil, errors.New("privatekey not found in config")
	}

	// Tạo đối tượng Token với Claims và SigningMethod
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký chuỗi token với khóa
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("error signing token: %v", err)
	}

	return &tokenString, nil
}

func GenerateRefreshToken(email string) (*string, error) {

	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * (7 * 24)).Unix(),
	}
	privateKey := os.Getenv("PRIVATE_KEY")
	fmt.Println("Private Key:", privateKey)
	if privateKey == "" {
		log.Info(" privateKey == null")
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString([]byte(privateKey))
	if err != nil {
		log.Error(err)
	}

	return &refresh, nil

}

func extractToken(authorizationHeader string) string {
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	log.Info("Invalid token format. Parts:", parts)
	return ""
}
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := extractToken(authorizationHeader)
		//revokeExpiredToken := RevokeExpiredToken(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		c.Next()
	}
}
