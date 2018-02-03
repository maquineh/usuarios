package routers

import (
	"github.com/labstack/echo"
	"github.com/maquineh/usuarios/controllers"
)

//App é uma instância de App
var App *echo.Echo

func init() {
	App = echo.New()
	App.GET("/", controllers.Home)
}
