package domain

import (
	"fmt"
	"strconv"

	"github.com/kartesus/bank.go/internal/platform"
)

type ProcessDepositPresenter interface {
	AccountNotFound(id string)
	InvalidParam(paramName string, paramValue any, reason string)
	DepositProcessed(account map[string]any)
}

type ProcessDepositHandler struct {
	store platform.Store
}

func NewProcessDepositHandler(store platform.Store) *ProcessDepositHandler {
	return &ProcessDepositHandler{store: store}
}

func (h *ProcessDepositHandler) Handle(req map[string]string, p ProcessDepositPresenter) {
	id := req["id"]

	amount, err := strconv.ParseInt(req["amount"], 10, 64)

	if err != nil {
		p.InvalidParam("amount", req["amount"], "must be a number")
		return
	}

	if id == "" {
		p.InvalidParam("id", id, "must not be empty")
		return
	}

	if amount <= 0 {
		p.InvalidParam("amount", amount, "must be greater than zero")
		return
	}

	account, ok := h.store.Get(id)

	if !ok {
		p.AccountNotFound(id)
		return
	}

	fmt.Printf("account: %v", account)

	bonus := amount / 200
	account["balance"] = account["balance"].(int64) + amount + bonus
	account["transactions"] = append(account["transactions"].([]map[string]any), map[string]any{"amount": amount + bonus, "type": "deposit"})

	p.DepositProcessed(account)
}
