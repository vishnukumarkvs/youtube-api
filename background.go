package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vishnukumarkvs/youtube-api/config"
	"github.com/vishnukumarkvs/youtube-api/database"
	"github.com/vishnukumarkvs/youtube-api/models"
)

type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			PublishedAt time.Time `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

func FetchVideos(cfg config.Config) {
	log.Println("Fetching video data started")
	for {
		log.Println("Fetch in loop started")
		resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&order=date&type=video&q=official&key=%s", cfg.APIKey))
		if err != nil {
			log.Println("Error fetching YouTube API:", err)
			continue
		}
		defer resp.Body.Close()

		log.Printf("API Response Status: %s\n", resp.Status)

		var ytResponse YouTubeResponse
		if err := json.NewDecoder(resp.Body).Decode(&ytResponse); err != nil {
			log.Println("Error decoding YouTube response:", err)
			continue
		}

		log.Printf("Number of items received: %d\n", len(ytResponse.Items))


		for _, item := range ytResponse.Items {
			video := models.Video{
				ID:          item.ID.VideoID,
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: item.Snippet.PublishedAt,
				Thumbnail:   item.Snippet.Thumbnails.Default.URL,
			}
			log.Println("Video:", video.Title)
			_, err := database.DB.NamedExec(`INSERT INTO videos (id, title, description, published_at, thumbnail) VALUES (:id, :title, :description, :published_at, :thumbnail) ON CONFLICT (id) DO NOTHING`, &video)
			if err != nil {
				log.Println("Error inserting video into database:", err)
			}else{
				log.Println("Video inserted into database:", video.Title)
			}
		}

		time.Sleep(10 * time.Second)
	}
}
