package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWTToken(username string, expiration time.Duration) (string, error) {
	claims := JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func CustomJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.String(http.StatusUnauthorized, "Invalid token format")
		}

		token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.String(http.StatusUnauthorized, "Invalid token claims")
		}

		username, ok := claims["username"].(string)
		if !ok {
			return c.String(http.StatusUnauthorized, "Invalid Username")
		}

		c.Set("username", username)
		return next(c)
	}
}
