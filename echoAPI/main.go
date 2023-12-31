package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// e can be used to access everything in/from echo
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!, Here I am")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
