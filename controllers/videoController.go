package controllers

import (
	"log"
	"net/http"
	"time"

	"my_tube_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DB is a global variable that holds the database connection
var DB *gorm.DB

func CreateVideo(c *gin.Context) {
	if DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	var video models.Video
	if err := c.BindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	video.TimeCreated = time.Now()

	if err := DB.Create(&video).Error; err != nil {
		// Log the specific error message for debugging
		log.Printf("Failed to create video: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

func GetVideos(c *gin.Context) {
	if DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	var videos []models.Video
	if err := DB.Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	c.JSON(http.StatusOK, videos)
}

func GetVideo(c *gin.Context) {
	if DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	id := c.Param("id")
	var video models.Video
	if err := DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	c.JSON(http.StatusOK, video)
}

func UpdateVideo(c *gin.Context) {
	if DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	id := c.Param("id")
	var existingVideo models.Video
	if err := DB.First(&existingVideo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	var updatedVideo models.Video
	if err := c.BindJSON(&updatedVideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update the existing video with the new data
	existingVideo.Title = updatedVideo.Title
	existingVideo.Category = updatedVideo.Category
	existingVideo.CreatedBy = updatedVideo.CreatedBy
	existingVideo.Description = updatedVideo.Description
	existingVideo.Views = updatedVideo.Views
	existingVideo.Comments = updatedVideo.Comments
	existingVideo.Subscribers = updatedVideo.Subscribers

	if err := DB.Save(&existingVideo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}

	c.JSON(http.StatusOK, existingVideo)
}

func DeleteVideo(c *gin.Context) {
	if DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	id := c.Param("id")
	if err := DB.Delete(&models.Video{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
}
