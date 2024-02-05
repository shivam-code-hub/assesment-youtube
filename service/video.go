package service

import (
	"log"
	"time"

	"google.golang.org/api/youtube/v3"

	"youtube/connection"
	"youtube/models"
)

func SaveVideoData(videos []*youtube.SearchResult) {
	db := connection.ConnectDatabase()

	for _, item := range videos {
		publishedAtTime, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			log.Printf("Error parsing PublishedAt time: %v", err)
			continue // Skip if the time can't be parsed
		}

		video := models.Video{
			VideoID:            item.Id.VideoId,
			Title:              item.Snippet.Title,
			Description:        item.Snippet.Description,
			PublishingDatetime: publishedAtTime,
			ThumbnailURL:       item.Snippet.Thumbnails.Default.Url,
		}

		// Insert or update the record
		err = db.Where(models.Video{VideoID: video.VideoID}).Assign(video).FirstOrCreate(&video).Error
		if err != nil {
			log.Printf("Error saving video data: %v", err)
		}
	}
}
