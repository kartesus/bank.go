package domain

import "github.com/kartesus/bank.go/internal/platform"

type ForHandlingListAccountsResult interface {
	ListAccounts(accounts []map[string]any)
}

type ListAccountsHandler struct {
	store platform.Store
}

func NewListAccountsHandler(store platform.Store) *ListAccountsHandler {
	return &ListAccountsHandler{store: store}
}

func (h *ListAccountsHandler) Handle(req map[string]string, res ForHandlingListAccountsResult) {
	accounts := h.store.GetAll()

	res.ListAccounts(accounts)
}