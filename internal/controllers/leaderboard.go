package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/coding-relay-be/internal/services"
	"github.com/siam-vit/coding-relay-be/internal/utils"
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

func StartTimer(c echo.Context) error {
	var input struct {
		Timer float64 `json:"timer"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	utils.CreateTimer(time.Hour * time.Duration(input.Timer))

	log.Println("Timer started")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully started timer",
	})
}

func GetTimeLeft(c echo.Context) error {
	if utils.GlobalTimer == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Timer has not been started yet",
		})
	}

	remainingTime := utils.GlobalTimer.TimeLeft()

	return c.JSON(http.StatusOK, map[string]int{
		"time_left": remainingTime / 1000000000,
	})
}
