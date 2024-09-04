package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

// MpesaCallback handles the callback from the M-Pesa API.
func MpesaCallback(ctx *gin.Context) {
	var callbackResponse map[string]interface{}

	// Bind the JSON body to the callbackResponse map
	if err := ctx.ShouldBindJSON(&callbackResponse); err != nil {
		log.Printf("Error parsing callback JSON: %v", err)
		ctx.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Process the callback response
	log.Printf("Mpesa Callback Data: %+v", callbackResponse)

	// Send a response indicating receipt of the callback
	ctx.JSON(200, gin.H{"status": "Callback received"})
}
