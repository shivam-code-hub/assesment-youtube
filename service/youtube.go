package service

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchLatestVideos(searchQuery string) ([]*youtube.SearchResult, error) {
	ctx := context.Background()

	apiKey := GetNextAPIKey()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	call := service.Search.List([]string{"id", "snippet"}).
		Q(searchQuery).
		MaxResults(10).
		Order("date")
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	return response.Items, nil
}
