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

// Register handles user registration, creates a profile, and generates a JWT token
func (ac *AuthController) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password before storing it
    if err := user.SetPassword(user.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user.Role = "user" // Default role

    if err := ac.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    // Create a profile for the new user
    profile := models.Profile{
        UserID: user.ID,
        // Set other profile fields if needed
    }

    if err := ac.DB.Create(&profile).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
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
        "profile": profile,
        "token":   token,
    })
}

// Login handles user login, retrieves profile, and returns user details along with a JWT token
func (ac *AuthController) Login(c *gin.Context) {
    var loginUser models.User
    var user models.User

    if err := c.ShouldBindJSON(&loginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find the user by email
    if err := ac.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
        // Return unauthorized if user is not found
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Check if the password matches
    if err := user.CheckPassword(loginUser.Password); err != nil {
        // Return unauthorized if password is incorrect
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Retrieve the user's profile
    var profile models.Profile
    if err := ac.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Profile not found, create a new empty profile
            profile = models.Profile{UserID: user.ID}
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    // Generate JWT token after successful login
    token, err := utils.GenerateJWT(user.Email, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return 
    }

    c.JSON(http.StatusOK, gin.H{
        "user":    user,
        "profile": profile,
        "token":   token,
    })
}
