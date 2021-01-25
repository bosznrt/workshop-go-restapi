package user

import (
	"server/db"
	"server/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// UserRouter Create user router
func UserRouter(app *echo.Group, resource *db.Resource) {
	repository := NewUserRepository(resource)

	decrypt := middleware.JWTConfig{
		Claims:     &utils.MyCustomClaims{},
		SigningKey: []byte("something you want"),
	}

	app.POST("/users/register", handleRegisterUser(repository))
	app.GET("/users", handleGetAllUsers(repository), middleware.JWTWithConfig(decrypt))
	app.GET("/users/:id", handleGetUserById(repository), middleware.JWTWithConfig(decrypt))
}
