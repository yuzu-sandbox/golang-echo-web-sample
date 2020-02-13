package main

import (
	"echo-sample/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/user/create", controllers.CreateUser)
}
