package controllers

import (
	"github.com/gin-gonic/gin"
	"my_tube_backend/services"
)

// MpesaController handles M-Pesa related operations.
type MpesaController struct {
	Service *services.MpesaService
}

// NewMpesaController creates a new instance of MpesaController.
func NewMpesaController(service *services.MpesaService) *MpesaController {
	return &MpesaController{Service: service}
}

// STKPush handles the STK push request.
func (c *MpesaController) STKPush(ctx *gin.Context) {
	var stkRequest struct {
		PhoneNumber      string `json:"phoneNumber"`
		Amount           string `json:"amount"`
		AccountReference string `json:"accountReference,omitempty"`
		TransactionDesc  string `json:"transactionDesc,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&stkRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if stkRequest.AccountReference == "" {
		stkRequest.AccountReference = "DefaultAccountReference"
	}
	if stkRequest.TransactionDesc == "" {
		stkRequest.TransactionDesc = "DefaultTransactionDescription"
	}

	response, err := c.Service.STKPush(stkRequest.PhoneNumber, stkRequest.Amount, stkRequest.AccountReference, stkRequest.TransactionDesc)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}
