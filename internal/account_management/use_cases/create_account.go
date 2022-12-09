package domain

import (
	"github.com/kartesus/bank.go/internal/platform"
)

type CreateAccountPresenter interface {
	AccountAlreadyExists(id string)
	AccountCreated(account map[string]any)
	InvalidParam(paramName string, paramValue any, reason string)
}

type CreateAccountHandler struct {
	store platform.Store
}

func NewCreateAccountHandler(store platform.Store) *CreateAccountHandler {
	return &CreateAccountHandler{store: store}
}

func (h *CreateAccountHandler) Handle(req map[string]string, res CreateAccountPresenter) {
	id := req["id"]

	if id == "" {
		res.InvalidParam("id", id, "must not be empty")
		return
	}

	if h.store.HasKey(id) {
		res.AccountAlreadyExists(id)
		return
	}

	customerName := req["customerName"]

	if customerName == "" {
		res.InvalidParam("customerName", customerName, "must not be empty")
		return
	}

	fiscalDocument := req["fiscalDocument"]

	if fiscalDocument == "" {
		res.InvalidParam("fiscalDocument", fiscalDocument, "must not be empty")
		return
	}

	entity := map[string]any{
		"id":             id,
		"customerName":   customerName,
		"fiscalDocument": fiscalDocument,
		"balance":        int64(0),
		"transactions":   []map[string]any{},
	}

	h.store.Put(id, entity)
	res.AccountCreated(entity)
}
