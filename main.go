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
	"my_tube_backend/controllers" 
	"my_tube_backend/services"    // Import your M-Pesa services
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
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
	if err := db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Video{}, &models.Comment{}, &models.Subscriber{}); err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Log all incoming requests
	r.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Assign DB instance to the controllers package
	controllers.DB = db

	// Initialize M-Pesa Service
	mpesaService, err := services.NewMpesaService()
	if err != nil {
		log.Fatal("Error initializing Mpesa Service: ", err)
	}

	// Initialize M-Pesa Controller
	mpesaController := controllers.NewMpesaController(mpesaService)

	// Setup routes for authentication and user management
	routes.AuthRoutes(r, db)

	// Setup routes for file handling
	r.POST("/upload", controllers.UploadFileController)
	r.POST("/resize", controllers.ResizeImageController)

	// Setup routes for video management
	routes.RegisterVideoRoutes(r)

	// Setup additional routes
	routes.SetupRoutes(db, r)

	// Setup M-Pesa routes
	routes.MpesaRoutes(r, mpesaController)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
