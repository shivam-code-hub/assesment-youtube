package route

import (
	"github.com/gin-gonic/gin"

	"youtube/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/videos", controller.GetVideos)
	router.GET("/search", controller.SearchVideos)

	return router
}
