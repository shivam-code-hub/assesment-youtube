package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"youtube/connection"
	"youtube/models"
	"youtube/route"
	"youtube/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := connection.ConnectDatabase()

	err = db.AutoMigrate(&models.Video{})
	if err != nil {
		log.Fatal("Error auto migrating database: ", err)
	}
}

func main() {
	go func() {
		ticker := time.NewTicker(30 * time.Second) // Fetch videos every 30 seconds
		for range ticker.C {
			searchQuery := os.Getenv("YOUTUBE_SEARCH_QUERY")
			if searchQuery == "" {
				log.Println("YOUTUBE_SEARCH_QUERY is not set in .env file")
				continue
			}
			videos, err := service.FetchLatestVideos(searchQuery)
			if err != nil {
				log.Println("Error fetching videos: ", err)
				continue
			}
			service.SaveVideoData(videos)
		}
	}()

	router := route.SetupRouter()
	router.Run(":8080")
}
