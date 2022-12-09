package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kartesus/bank.go/cmd/api"
	"github.com/kartesus/bank.go/internal/platform"
)

func main() {
	platform := platform.NewTestPlatform()
	accountManagement := api.NewAccountManagementController(platform)

	router := gin.Default()
	router.POST("/accounts", accountManagement.CreateAccount)
	router.GET("/accounts/:id", accountManagement.RetrieveAccount)
	router.Run(":8080")
}
