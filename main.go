package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vishnukumarkvs/youtube-api/config"
	"github.com/vishnukumarkvs/youtube-api/database"
	"github.com/vishnukumarkvs/youtube-api/handlers"
)

func main() {
	cfg := config.GetConfig()

	// Connect to the database
	database.Connect(cfg)

	// Start background task
	go FetchVideos(cfg)

	// Set up Echo
	e := echo.New()

	e.GET("/videos", handlers.GetVideos)
	e.GET("/search", handlers.SearchVideos)
	e.GET("/", handlers.HealthCheck)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
