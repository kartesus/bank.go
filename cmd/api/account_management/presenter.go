package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountManagementPresenter struct {
	ctx *gin.Context
}

func (h *accountManagementPresenter) AccountAlreadyExists(id string) {
	h.ctx.JSON(http.StatusBadRequest, gin.H{"error": "account already exists"})
}

func (h *accountManagementPresenter) AccountCreated(account map[string]any) {
	h.ctx.JSON(http.StatusCreated, account)
}

func (h *accountManagementPresenter) InvalidParam(paramName string, paramValue any, reason string) {
	h.ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue.(string) + ", but " + reason})
}

func (h *accountManagementPresenter) AccountNotFound(id string) {
	h.ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}

func (h *accountManagementPresenter) DepositProcessed(account map[string]any) {
	h.ctx.JSON(http.StatusOK, account)
}

func (h *accountManagementPresenter) TransferProcessed(fromAccount map[string]any, toAccount map[string]any, amount int64) {
	h.ctx.JSON(http.StatusOK, map[string]any{"fromAccount": fromAccount, "toAccount": toAccount, "amount": amount})
}

func (h *accountManagementPresenter) InsufficientFunds(id string) {
	h.ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "insufficient funds"})
}

func (h *accountManagementPresenter) TransferToSameAccount(id string) {
	h.ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "transfer to same account"})
}

func (h *accountManagementPresenter) TransferToNonExistingAccount(id string) {
	h.ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "transfer to non-existing account"})
}

func (h *accountManagementPresenter) WithdrawProcessed(account map[string]any) {
	h.ctx.JSON(http.StatusOK, account)
}

func (h *accountManagementPresenter) AccountRetrieved(account map[string]any) {
	h.ctx.JSON(http.StatusOK, account)
}

func (h *accountManagementPresenter) ListAccounts(accounts []map[string]any) {
	h.ctx.JSON(http.StatusOK, accounts)
}
