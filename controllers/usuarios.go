package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Página inicial da aplicação
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
