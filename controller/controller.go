package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"youtube/connection"
	"youtube/models"
)

func GetVideos(c *gin.Context) {
	var videos []models.Video
	page, _ := c.GetQuery("page")
	limit, _ := c.GetQuery("limit")

	pageNumber, _ := strconv.Atoi(page)
	limitNumber, _ := strconv.Atoi(limit)

	if pageNumber < 1 {
		pageNumber = 1
	}
	if limitNumber <= 0 {
		limitNumber = 10
	}

	offset := (pageNumber - 1) * limitNumber

	db := connection.ConnectDatabase()

	res := db.Order("publishing_datetime desc").Limit(limitNumber).Offset(offset).Find(&videos)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": videos})
}

func SearchVideos(c *gin.Context) {
	var videos []models.Video
	searchQuery := c.Query("query")

	db := connection.ConnectDatabase()

	res := db.Debug().Where("title LIKE ? OR description LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%").Find(&videos)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": videos})
}
