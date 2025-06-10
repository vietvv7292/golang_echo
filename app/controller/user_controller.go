package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"Alice", "Bob"})
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Created user: " + name,
	})
}
