package main

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/kirimi/rb2backend/hello"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet("World"))
	})
	_ = e.Start(":8081")
}
