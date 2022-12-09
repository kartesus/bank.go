package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type retrieveAccountResponseHandler struct {
	c *gin.Context
}

func (h *retrieveAccountResponseHandler) AccountNotFound(id string) {
	h.c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}

func (h *retrieveAccountResponseHandler) AccountRetrieved(account map[string]any) {
	h.c.JSON(http.StatusOK, account)
}

func (h *retrieveAccountResponseHandler) InvalidParam(paramName string, paramValue string, reason string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue + ", but " + reason})
}
