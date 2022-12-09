package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type processWithdrawResponseHandler struct {
	c *gin.Context
}

func (h *processWithdrawResponseHandler) AccountNotFound(id string) {
	h.c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}

func (h *processWithdrawResponseHandler) WithdrawProcessed(account map[string]any) {
	h.c.JSON(http.StatusOK, account)
}

func (h *processWithdrawResponseHandler) InvalidParam(paramName string, paramValue any, reason string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue.(string) + ", but " + reason})
}

func (h *processWithdrawResponseHandler) InsufficientFunds(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "insufficient funds"})
}
