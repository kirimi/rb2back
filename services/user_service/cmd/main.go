package main

import (
	"fmt"
	"github.com/kirimi/rb2backend/user_service/internal/config"
	"github.com/labstack/echo"
	"net/http"

	"github.com/kirimi/rb2backend/hello"
)

func main() {
	cfg := config.GetConfig("config.yaml")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SslMode,
	)

	fmt.Println(dsn)

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet(dsn))
	})
	_ = e.Start(":8080")
}
