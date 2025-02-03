package main

import (
	"net/http"

    "github.com/siam-vit/coding-relay-be/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Echo!"})
	})
	e.Start(":8080")
}
