package auth

import (
	"net/http"
	"server/user"
	"server/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func handleLogin(repository user.Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		user, err := repository.VerifyUser(email, password)

		if user == nil || err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		claims := &utils.MyCustomClaims{
			user.Id.String(),
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		}

		mySigningKey := []byte("something you want")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		issued, err := token.SignedString(mySigningKey)

		return c.JSON(http.StatusOK, echo.Map{
			"token": issued,
		})
	}
}
