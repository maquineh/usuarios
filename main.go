package main

import (
	"github.com/labstack/echo/middleware"
	r "github.com/maquineh/usuarios/routers"
)

func main() {
	e := r.App
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))
}
