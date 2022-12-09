package platform

import "github.com/kartesus/bank.go/internal/account_management/infrastructure"

type Platform struct {
	AccountStore Store
}

func NewTestPlatform() *Platform {
	return &Platform{AccountStore: infrastructure.NewMemoryStore()}
}
