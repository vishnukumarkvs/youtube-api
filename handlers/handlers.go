package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vishnukumarkvs/youtube-api/database"
	"github.com/vishnukumarkvs/youtube-api/models"
)

func GetVideos(c echo.Context) error {
	var videos []models.Video
	err := database.DB.Select(&videos, "SELECT * FROM videos ORDER BY published_at DESC LIMIT 10")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, videos)
}

func SearchVideos(c echo.Context) error {
	query := c.QueryParam("q")
	var videos []models.Video
	err := database.DB.Select(&videos, "SELECT * FROM videos WHERE title ILIKE $1 OR description ILIKE $1 ORDER BY published_at DESC", "%"+query+"%")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, videos)
}

func HealthCheck(c echo.Context) error {
	log.Println("App is running in port 8080")
	return c.String(http.StatusOK, "App is Running")
}