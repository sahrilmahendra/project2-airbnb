package middlewares

import (
	"project2/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// function untuk membuat token
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() // expired dalam 5 jam

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))

}

// function untuk mendapatkan user id
func ExtractTokenId(e echo.Context) int {
	users := e.Get("user").(*jwt.Token)
	if users.Valid {
		claims := users.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
