package routes

import (
	"github.com/gin-gonic/gin"
	"my_tube_backend/controllers"
)

// RegisterVideoRoutes sets up routes for video management
func RegisterVideoRoutes(router *gin.Engine) {
	videoRoutes := router.Group("/videos")
	{
		videoRoutes.POST("/upload", controllers.CreateVideo)
		videoRoutes.GET("/", controllers.GetVideos)
		videoRoutes.GET("/:id", controllers.GetVideo)
		videoRoutes.PUT("/:id", controllers.UpdateVideo)
		videoRoutes.DELETE("/:id", controllers.DeleteVideo)
	}
}
