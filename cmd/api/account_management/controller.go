package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/kartesus/bank.go/internal/account_management/use_cases"
	"github.com/kartesus/bank.go/internal/platform"
)

type accountManagementController struct {
	createAccount   *uc.CreateAccountHandler
	retrieveAccount *uc.RetrieveAccountHandler
	processDeposit  *uc.ProcessDepositHandler
	processWithdraw *uc.ProcessWithdrawHandler
	processTransfer *uc.ProcessTransferHandler
	listAccounts    *uc.ListAccountsHandler
}

func NewAccountManagementController(platform *platform.Platform) *accountManagementController {
	return &accountManagementController{
		createAccount:   uc.NewCreateAccountHandler(platform.AccountStore),
		retrieveAccount: uc.NewRetrieveAccountHandler(platform.AccountStore),
		processDeposit:  uc.NewProcessDepositHandler(platform.AccountStore),
		processWithdraw: uc.NewProcessWithdrawHandler(platform.AccountStore),
		processTransfer: uc.NewProcessTransferHandler(platform.AccountStore),
		listAccounts:    uc.NewListAccountsHandler(platform.AccountStore),
	}
}

func (api accountManagementController) CreateAccount(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := &accountManagementPresenter{c}
	api.createAccount.Handle(req, res)
}

func (api accountManagementController) RetrieveAccount(c *gin.Context) {
	id := c.Param("id")

	res := &accountManagementPresenter{c}
	api.retrieveAccount.Handle(map[string]string{"id": id}, res)
}

func (api accountManagementController) ProcessDeposit(c *gin.Context) {
	id := c.Param("id")

	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req["id"] = id

	res := &accountManagementPresenter{c}
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

	res := &accountManagementPresenter{c}
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

	res := &accountManagementPresenter{c}
	api.processTransfer.Handle(req, res)
}

func (api accountManagementController) ListAccounts(c *gin.Context) {
	res := &accountManagementPresenter{c}
	api.listAccounts.Handle(map[string]string{}, res)
}
