package handler

import (
	"net/http"
	"startup-crowdfunding-backend/helper"
	"startup-crowdfunding-backend/transaction"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetTransactionsByCampaignIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get transactions failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get transactions failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatCampaignTransactions(transactions)
	response := helper.APIResponse("Get transactions success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
