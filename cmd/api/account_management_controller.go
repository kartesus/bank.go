package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kartesus/bank.go/internal/account_management/domain"
	"github.com/kartesus/bank.go/internal/platform"
)

type accountManagementController struct {
	createAccount *domain.CreateAccountHandler
}

func NewAccountManagementController(platform *platform.Platform) *accountManagementController {
	return &accountManagementController{createAccount: domain.NewCreateAccountHandler(platform.AccountStore)}
}

func (api accountManagementController) CreateAccount(c *gin.Context) {
	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := &createAccountResponseHandler{c}
	api.createAccount.Handle(req, res)
}
