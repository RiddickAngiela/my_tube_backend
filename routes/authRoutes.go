package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "my_tube_backend/controllers"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
    authController := &controllers.AuthController{DB: db}

    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/register", authController.Register)
        authRoutes.POST("/login", authController.Login)
    }
}
