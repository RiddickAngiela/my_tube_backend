package routes

import (
	"github.com/gin-gonic/gin"
	"my_tube_backend/controllers"
)

// MpesaRoutes sets up the routes for M-Pesa related endpoints.
func MpesaRoutes(r *gin.Engine, mpesaController *controllers.MpesaController) {
	r.POST("/mpesa/stkpush", mpesaController.STKPush)
	r.POST("/mpesa/callback", controllers.MpesaCallback)
}
