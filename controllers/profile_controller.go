package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "my_tube_backend/models"
)

// CreateProfile handles the creation of a new profile using Gin and GORM
func CreateProfile(c *gin.Context, db *gorm.DB) {
    var profile models.Profile
    if err := c.ShouldBindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&profile).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, profile)
}

// GetProfile handles retrieving a profile by ID using Gin and GORM
func GetProfile(c *gin.Context, db *gorm.DB) {
    id := c.Param("id")
    var profile models.Profile

    if err := db.First(&profile, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, profile)
}
