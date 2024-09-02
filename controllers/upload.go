package controllers

import (
	"net/http"
	"my_tube_backend/utils"
	"github.com/gin-gonic/gin"
)

// UploadFileController handles file uploads.
func UploadFileController(c *gin.Context) {
	err := utils.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
