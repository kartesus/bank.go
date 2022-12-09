package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type processTransferResponseHandler struct {
	c *gin.Context
}

func (h *processTransferResponseHandler) AccountNotFound(id string) {
	h.c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}

func (h *processTransferResponseHandler) TransferProcessed(fromAccount map[string]any, toAccount map[string]any, amount int64) {
	h.c.JSON(http.StatusOK, map[string]any{"fromAccount": fromAccount, "toAccount": toAccount, "amount": amount})
}

func (h *processTransferResponseHandler) InvalidParam(paramName string, paramValue any, reason string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue.(string) + ", but " + reason})
}

func (h *processTransferResponseHandler) InsufficientFunds(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "insufficient funds"})
}

func (h *processTransferResponseHandler) TransferToSameAccount(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "transfer to same account"})
}

func (h *processTransferResponseHandler) TransferToNonExistingAccount(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "transfer to non-existing account"})
}
