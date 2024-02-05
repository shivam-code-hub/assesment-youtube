package models

import "time"

type Video struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	VideoID            string    `gorm:"uniqueIndex;type:varchar(255)" json:"video_id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	PublishingDatetime time.Time `json:"publishing_datetime"`
	ThumbnailURL       string    `json:"thumbnail_url"`
}
