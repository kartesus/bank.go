package domain

import "github.com/kartesus/bank.go/internal/platform"

type RetrieveAccountPresenter interface {
	AccountNotFound(id string)
	AccountRetrieved(account map[string]interface{})
	InvalidParam(paramName string, paramValue any, reason string)
}

type RetrieveAccountHandler struct {
	store platform.Store
}

func NewRetrieveAccountHandler(store platform.Store) *RetrieveAccountHandler {
	return &RetrieveAccountHandler{store: store}
}

func (h *RetrieveAccountHandler) Handle(req map[string]string, p RetrieveAccountPresenter) {
	id := req["id"]

	if id == "" {
		p.InvalidParam("id", id, "must not be empty")
		return
	}

	account, ok := h.store.Get(id)

	if !ok {
		p.AccountNotFound(id)
		return
	}

	p.AccountRetrieved(account)
}
