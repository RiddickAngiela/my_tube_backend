package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "my_tube_backend/controllers"
)

// SetupRoutes initializes the routes for the application using Gin and GORM
func SetupRoutes(db *gorm.DB, r *gin.Engine) {
    // Profile routes
    r.POST("/profile", func(c *gin.Context) {
        controllers.CreateProfile(c, db)
    })
    r.GET("/profile/:id", func(c *gin.Context) {
        controllers.GetProfile(c, db)
    })
}
