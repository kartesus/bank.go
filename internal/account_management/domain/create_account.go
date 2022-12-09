package domain

import (
	"github.com/kartesus/bank.go/internal/platform"
)

type ForHandlingCreateAccountResult interface {
	AccountAlreadyExists(id string)
	AccountCreated(account map[string]any)
	InvalidParam(paramName string, paramValue string, reason string)
}

type CreateAccountHandler struct {
	store platform.Store
}

func NewCreateAccountHandler(store platform.Store) *CreateAccountHandler {
	return &CreateAccountHandler{store: store}
}

func (h *CreateAccountHandler) Handle(req map[string]interface{}, res ForHandlingCreateAccountResult) {
	id := req["id"].(string)

	if id == "" {
		res.InvalidParam("id", id, "must not be empty")
		return
	}

	if h.store.HasKey(id) {
		res.AccountAlreadyExists(id)
		return
	}

	customerName := req["customerName"].(string)

	if customerName == "" {
		res.InvalidParam("customerName", customerName, "must not be empty")
		return
	}

	fiscalDocument := req["fiscalDocument"].(string)

	if fiscalDocument == "" {
		res.InvalidParam("fiscalDocument", fiscalDocument, "must not be empty")
		return
	}

	entity := map[string]any{"id": id, "customerName": customerName, "fiscalDocument": fiscalDocument}
	h.store.Put(id, entity)
	res.AccountCreated(entity)
}
