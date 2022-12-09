package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kartesus/bank.go/internal/account_management/domain"
	"github.com/kartesus/bank.go/internal/platform"
)

type accountManagementController struct {
	createAccount   *domain.CreateAccountHandler
	retrieveAccount *domain.RetrieveAccountHandler
}

func NewAccountManagementController(platform *platform.Platform) *accountManagementController {
	return &accountManagementController{
		createAccount:   domain.NewCreateAccountHandler(platform.AccountStore),
		retrieveAccount: domain.NewRetrieveAccountHandler(platform.AccountStore),
	}
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

func (api accountManagementController) RetrieveAccount(c *gin.Context) {
	id := c.Param("id")

	res := &retrieveAccountResponseHandler{c}
	api.retrieveAccount.Handle(map[string]any{"id": id}, res)
}
