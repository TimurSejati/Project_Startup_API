package handler

import (
	"net/http"
	"startup/helper"
	"startup/transaction"
	"startup/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's of transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.transactionService.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's of transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Transactions detail", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.transactionService.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's of transactionsa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
