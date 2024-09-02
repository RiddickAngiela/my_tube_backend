package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"my_tube_backend/utils"
)

func ResizeImageController(c *gin.Context) {
	err := utils.ResizeImage(c.Request.Body, 100, 100) // Adjust dimensions as needed
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image resized successfully"})
}
