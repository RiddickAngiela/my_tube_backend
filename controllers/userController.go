package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "my_tube_backend/models"
)

type UserController struct {
    DB *gorm.DB
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
    var users []models.User
    if err := uc.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
    var user models.User
    if err := uc.DB.Where("id = ?", c.PostForm("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := uc.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
    var user models.User
    if err := uc.DB.Where("id = ?", c.PostForm("id")).Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
