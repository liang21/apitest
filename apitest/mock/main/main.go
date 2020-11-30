package main

import (
	"apitest/apitest/mock/mock"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	registerRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func registerRouter(e *echo.Echo) {
	e.POST("/api/users", mock.CreateUser)
	e.PUT("/api/users/:name", mock.UpdateUser)
	e.GET("/api/users/:name", mock.GetUser)
	e.DELETE("/api/users/:name", mock.DeleteUser)
}
