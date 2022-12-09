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
	processDeposit  *domain.ProcessDepositHandler
	processWithdraw *domain.ProcessWithdrawHandler
	processTransfer *domain.ProcessTransferHandler
}

func NewAccountManagementController(platform *platform.Platform) *accountManagementController {
	return &accountManagementController{
		createAccount:   domain.NewCreateAccountHandler(platform.AccountStore),
		retrieveAccount: domain.NewRetrieveAccountHandler(platform.AccountStore),
		processDeposit:  domain.NewProcessDepositHandler(platform.AccountStore),
		processWithdraw: domain.NewProcessWithdrawHandler(platform.AccountStore),
		processTransfer: domain.NewProcessTransferHandler(platform.AccountStore),
	}
}

func (api accountManagementController) CreateAccount(c *gin.Context) {
	var req map[string]string
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

func (api accountManagementController) ProcessDeposit(c *gin.Context) {
	id := c.Param("id")

	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req["id"] = id

	res := &processDepositResponseHandler{c}
	api.processDeposit.Handle(req, res)
}

func (api accountManagementController) ProcessWithdraw(c *gin.Context) {
	id := c.Param("id")

	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req["id"] = id

	res := &processWithdrawResponseHandler{c}
	api.processWithdraw.Handle(req, res)
}

func (api accountManagementController) ProcessTransfer(c *gin.Context) {
	id := c.Param("id")

	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req["originId"] = id

	res := &processTransferResponseHandler{c}
	api.processTransfer.Handle(req, res)
}
