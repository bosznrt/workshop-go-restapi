package main

import (
	"net/http"
	"server/auth"
	"server/db"
	"server/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Validator is implementation of validation of rquest values.
type Validator struct {
	validator *validator.Validate
}

// Validate do validation for request value.
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func main() {
	app := echo.New()
	validate := validator.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", healthcheck)
	app.Validator = &Validator{validator: validate}

	v1 := app.Group("/api/v1")

	resource, err := db.Init()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	user.UserRouter(v1, resource)
	auth.AuthRouter(v1, resource)

	app.Start(":8080")
}
