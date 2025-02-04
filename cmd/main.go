package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/siam-vit/coding-relay-be/internal/database"
	"github.com/siam-vit/coding-relay-be/internal/routes"
)

func main() {
	database.Connect()
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	routes.TeamRoutes(e)
	e.Start(":8080")
}
