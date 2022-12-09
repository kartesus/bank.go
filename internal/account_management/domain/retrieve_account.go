package domain

import "github.com/kartesus/bank.go/internal/platform"

type ForHandlingRetrieveAccountResult interface {
	AccountNotFound(id string)
	AccountRetrieved(account map[string]interface{})
	InvalidParam(paramName string, paramValue string, reason string)
}

type RetrieveAccountHandler struct {
	store platform.Store
}

func NewRetrieveAccountHandler(store platform.Store) *RetrieveAccountHandler {
	return &RetrieveAccountHandler{store: store}
}

func (h *RetrieveAccountHandler) Handle(req map[string]interface{}, res ForHandlingRetrieveAccountResult) {
	id := req["id"].(string)

	if id == "" {
		res.InvalidParam("id", id, "must not be empty")
		return
	}

	account, ok := h.store.Get(id)

	if !ok {
		res.AccountNotFound(id)
		return
	}

	res.AccountRetrieved(account)
}
