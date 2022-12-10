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

func (h *CreateAccountHandler) Handle(req map[string]string, p CreateAccountPresenter) {
	id := req["id"]

	if id == "" {
		p.InvalidParam("id", id, "must not be empty")
		return
	}

	if h.store.HasKey(id) {
		p.AccountAlreadyExists(id)
		return
	}

	customerName := req["customerName"]

	if customerName == "" {
		p.InvalidParam("customerName", customerName, "must not be empty")
		return
	}

	fiscalDocument := req["fiscalDocument"]

	if fiscalDocument == "" {
		p.InvalidParam("fiscalDocument", fiscalDocument, "must not be empty")
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
	p.AccountCreated(entity)
}
