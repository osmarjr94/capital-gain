package main

import (
	"echo"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func RegisterRoutes(e *echo.Echo) {
	repo := NewOperationRepository()
	service := NewOperationService(repo)
	controller := NewOperationController(service)

	e.POST("/operation", controller.HandleOperation)
}
