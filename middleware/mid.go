package middlewares

import (
	"errors"
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
func CreateTokenTeam(nip string, role string, userId int, department string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["nip"] = nip
	claims["role"] = role
	claims["userId"] = userId
	claims["department"] = department
	claims["exp"] = time.Now().Add(time.Duration(1) * time.Hour).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenTeamRole(e echo.Context) string {
	roles := e.Get("user").(*jwt.Token)
	if roles.Valid {
		claims := roles.Claims.(jwt.MapClaims)
		peran := claims["role"].(string)
		return peran
	}
	return ""
}
func ExtractTokenTeamNip(e echo.Context) string {
	nip := e.Get("user").(*jwt.Token)
	if nip.Valid {
		claims := nip.Claims.(jwt.MapClaims)
		usernip := claims["nip"].(string)
		return usernip
	}
	return ""
}
func ExtractTokenIdUser(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)

	if !user.Valid {
		return 0, errors.New("Invalid or expired token")
	}

	claims := user.Claims.(jwt.MapClaims)

	userId := int(claims["userId"].(float64))
	return userId, nil

}
func ExtractTokenTeamDepartment(e echo.Context) string {
	roles := e.Get("user").(*jwt.Token)
	if roles.Valid {
		claims := roles.Claims.(jwt.MapClaims)
		peran := claims["department"].(string)
		return peran
	}
	return ""
}
