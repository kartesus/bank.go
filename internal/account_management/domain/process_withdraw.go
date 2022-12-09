package domain

import (
	"strconv"

	"github.com/kartesus/bank.go/internal/platform"
)

type ForHandlingProcessWithdrawResult interface {
	AccountNotFound(id string)
	InvalidParam(paramName string, paramValue any, reason string)
	InsufficientFunds(id string)
	WithdrawProcessed(account map[string]any)
}

type ProcessWithdrawHandler struct {
	store platform.Store
}

func NewProcessWithdrawHandler(store platform.Store) *ProcessWithdrawHandler {
	return &ProcessWithdrawHandler{store: store}
}

func (h *ProcessWithdrawHandler) Handle(req map[string]string, res ForHandlingProcessWithdrawResult) {
	id := req["id"]

	amount, err := strconv.ParseInt(req["amount"], 10, 64)

	if err != nil {
		res.InvalidParam("amount", req["amount"], "must be a number")
		return
	}

	if id == "" {
		res.InvalidParam("id", id, "must not be empty")
		return
	}

	if amount <= 0 {
		res.InvalidParam("amount", amount, "must be greater than zero")
		return
	}

	account, ok := h.store.Get(id)

	if !ok {
		res.AccountNotFound(id)
		return
	}

	fee := amount / 100
	amount += fee

	if account["balance"].(int64) < amount {
		res.InsufficientFunds(id)
		return
	}

	account["balance"] = account["balance"].(int64) - amount
	account["transactions"] = append(account["transactions"].([]map[string]any), map[string]any{"amount": amount, "type": "withdraw"})

	res.WithdrawProcessed(account)
}
