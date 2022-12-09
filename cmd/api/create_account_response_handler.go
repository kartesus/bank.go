package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountResponseHandler struct {
	c *gin.Context
}

func (h *createAccountResponseHandler) AccountAlreadyExists(id string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "account already exists"})
}

func (h *createAccountResponseHandler) AccountCreated(account map[string]any) {
	h.c.JSON(http.StatusCreated, account)
}

func (h *createAccountResponseHandler) InvalidParam(paramName string, paramValue string, reason string) {
	h.c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param " + paramName + ". Received: " + paramValue + ", but " + reason})
}
