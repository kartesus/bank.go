package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountManagementResponseHandler struct {
	c *gin.Context
}

func (h *accountManagementResponseHandler) AccountAlreadyExists(id string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "account already exists"})
}

func (h *accountManagementResponseHandler) AccountCreated(account map[string]any) {
	h.c.JSON(http.StatusCreated, account)
}

func (h *accountManagementResponseHandler) InvalidParam(paramName string, paramValue any, reason string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue.(string) + ", but " + reason})
}

func (h *accountManagementResponseHandler) AccountNotFound(id string) {
	h.c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}

func (h *accountManagementResponseHandler) DepositProcessed(account map[string]any) {
	h.c.JSON(http.StatusOK, account)
}

func (h *accountManagementResponseHandler) TransferProcessed(fromAccount map[string]any, toAccount map[string]any, amount int64) {
	h.c.JSON(http.StatusOK, map[string]any{"fromAccount": fromAccount, "toAccount": toAccount, "amount": amount})
}

func (h *accountManagementResponseHandler) InsufficientFunds(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "insufficient funds"})
}

func (h *accountManagementResponseHandler) TransferToSameAccount(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "transfer to same account"})
}

func (h *accountManagementResponseHandler) TransferToNonExistingAccount(id string) {
	h.c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "transfer to non-existing account"})
}

func (h *accountManagementResponseHandler) WithdrawProcessed(account map[string]any) {
	h.c.JSON(http.StatusOK, account)
}

func (h *accountManagementResponseHandler) AccountRetrieved(account map[string]any) {
	h.c.JSON(http.StatusOK, account)
}
