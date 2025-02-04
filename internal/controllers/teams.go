package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/siam-vit/coding-relay-be/internal/models"
	"github.com/siam-vit/coding-relay-be/internal/services"
)

func CreateTeam(c echo.Context) error {

	var team models.Teams
	if err := c.Bind(&team); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create team",
			"data":    err.Error(),
		})
	}

	err := services.CreateTeam(team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create team",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully created team",
		"data":    team.TeamName,
	})

}
