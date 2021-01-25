package auth

import (
	"server/db"
	"server/user"

	"github.com/labstack/echo/v4"
)

func AuthRouter(app *echo.Group, resource *db.Resource) {
	repository := user.NewUserRepository(resource)

	app.POST("/login", handleLogin(repository))
}
