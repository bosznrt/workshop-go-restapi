package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserRequest ::model of user request

func handleRegisterUser(repository Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		code := http.StatusOK
		newUser := RegisterUser{}
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, getErrorMessage(err))
		}

		user, err := repository.AddUser(newUser)
		res := map[string]interface{}{
			"user": user,
			"err":  getErrorMessage(err),
		}

		return c.JSON(code, res)
	}
}

func handleGetAllUsers(repository Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		code := http.StatusOK
		users, err := repository.GetAllUsers()
		if err != nil {
			code = http.StatusInternalServerError
		}

		if len(users) == 0 {
			code = http.StatusNotFound
		}
		return c.JSON(code, users)
	}
}

func handleGetUserById(repository Repository) func(c echo.Context) error {
	return func(c echo.Context) error {

		id := c.Param("id")
		user, err := repository.GetUserByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, getErrorMessage(err))
		}

		response := map[string]interface{}{
			"user": user,
			"err":  getErrorMessage(err),
		}

		return c.JSON(http.StatusOK, response)
	}
}

func getErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
