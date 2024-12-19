package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/yourusername/cosmos-flash-loan/app/modules/flashloan/types"
)

// CreateFlashLoanRequest defines the structure for the flash loan request
type CreateFlashLoanRequest struct {
	Amount string `json:"amount" binding:"required"`
	Asset  string `json:"asset" binding:"required"`
}

// CreateFlashLoanHandler handles the creation of a flash loan
func CreateFlashLoanHandler(c *gin.Context) {
	var req CreateFlashLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest.ErrorResponse{Error: err.Error()})
		return
	}

	// Logic to process the flash loan request goes here
	// This would typically involve calling the keeper to handle the loan

	c.JSON(http.StatusOK, types.Response{Message: "Flash loan created successfully"})
}

// RegisterRoutes registers the REST routes for the flash loan module
func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/flashloan", CreateFlashLoanHandler)
}