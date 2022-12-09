package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type processDepositResponseHandler struct {
	c *gin.Context
}

func (h *processDepositResponseHandler) AccountNotFound(id string) {
	h.c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}

func (h *processDepositResponseHandler) DepositProcessed(account map[string]any) {
	h.c.JSON(http.StatusOK, account)
}

func (h *processDepositResponseHandler) InvalidParam(paramName string, paramValue any, reason string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue.(string) + ", but " + reason})
}
