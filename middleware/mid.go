package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
	})

}
func CreateTokenTeam(nip string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["nip"] = nip
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Duration(1) * time.Hour).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenTeamRole(e echo.Context) string {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		peran := claims["role"].(string)
		return peran
	}
	return ""
}
func ExtractTokenTeamId(e echo.Context) string {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		userId := claims["nip"].(string)
		return userId
	}
	return ""
}
