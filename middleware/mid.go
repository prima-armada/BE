package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateTokenTeam(teamId int, peran string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = teamId
	claims["peran"] = peran
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenTeamRole(e echo.Context) string {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		peran := claims["peran"].(string)
		return peran
	}
	return ""
}
func ExtractTokenTeamId(e echo.Context) int {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
