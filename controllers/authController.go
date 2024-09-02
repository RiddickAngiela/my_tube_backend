package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "my_tube_backend/models"
    "my_tube_backend/utils"
)

type AuthController struct {
    DB *gorm.DB
}

// Register handles user registration and generates a JWT token
func (ac *AuthController) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := user.SetPassword(user.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user.Role = "user" // Default role

    if err := ac.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    // Generate JWT token after successful registration
    token, err := utils.GenerateJWT(user.Email, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "user":    user,
        "token":   token,
    })
}

// Login handles user login and returns user details along with a JWT token
func (ac *AuthController) Login(c *gin.Context) {
    var loginUser models.User
    var user models.User

    if err := c.ShouldBindJSON(&loginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ac.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := user.CheckPassword(loginUser.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT token after successful login
    token, err := utils.GenerateJWT(user.Email, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return 
    }

    c.JSON(http.StatusOK, gin.H{
        "user":  user,
        "token": token,
    })
}