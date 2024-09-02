package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	"my_tube_backend/routes"
	"my_tube_backend/models"
	"my_tube_backend/controllers" // Import the controllers package
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Create the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database: ", err)
	} else {
		log.Println("Database connection established successfully.")
	}

	// AutoMigrate the models
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Log all incoming requests
	r.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Setup routes for authentication and user management
	routes.AuthRoutes(r, db)

	// Setup routes for file handling
	r.POST("/upload", controllers.UploadFileController)
	r.POST("/resize", controllers.ResizeImageController)

	// Run the server
	if err := r.Run(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
