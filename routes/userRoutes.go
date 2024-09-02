package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "my_tube_backend/controllers"
    "my_tube_backend/middlewares"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
    userController := &controllers.UserController{DB: db}

    userRoutes := router.Group("/users")
    userRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
    {
        userRoutes.POST("/get-all", userController.GetAllUsers)
        userRoutes.POST("/update", userController.UpdateUser)
        userRoutes.POST("/delete", userController.DeleteUser)
    }
}
