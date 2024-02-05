package service

import (
	"os"
	"sync/atomic"
)

var apiKeyIndex int32 = -1

func GetNextAPIKey() string {
	apiKeys := []string{
		os.Getenv("YOUTUBE_API_KEY_1"),
		os.Getenv("YOUTUBE_API_KEY_2"),
		os.Getenv("YOUTUBE_API_KEY_3"),
	}

	newIndex := atomic.AddInt32(&apiKeyIndex, 1) % int32(len(apiKeys))

	return apiKeys[newIndex]
}
