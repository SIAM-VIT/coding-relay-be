package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/coding-relay-be/internal/services"
)

func AddPoints(c echo.Context) error {
	id := c.Param("id")
	parseTeamID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse id",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	var points struct {
		Points uint `json:"points"`
	}
	if err := c.Bind(&points); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Incorrect input format",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	err = services.AddPoints(points.Points, parseTeamID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to add points",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully added points",
		"data":    "Successfully added points",
		"status":  "true",
	})
}

func ModifyPoints(c echo.Context) error {
	id := c.Param("id")
	parseTeamID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse id",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	var points struct {
		Points uint `json:"points"`
	}
	if err := c.Bind(&points); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Incorrect input format",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	err = services.ModifyPoints(points.Points, parseTeamID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to add points",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully added points",
		"data":    "Successfully added points",
		"status":  "true",
	})
}
