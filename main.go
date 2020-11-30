package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	name     string `json:"name"`
	password string `json:"password"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
