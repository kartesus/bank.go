package domain

import (
	"strconv"

	"github.com/kartesus/bank.go/internal/platform"
)

type ProcessTransferPresenter interface {
	AccountNotFound(id string)
	InvalidParam(paramName string, paramValue any, reason string)
	InsufficientFunds(id string)
	TransferToSameAccount(id string)
	TransferToNonExistingAccount(id string)
	TransferProcessed(originAccount map[string]any, destinationAccount map[string]any, amount int64)
}

type ProcessTransferHandler struct {
	store platform.Store
}

func NewProcessTransferHandler(store platform.Store) *ProcessTransferHandler {
	return &ProcessTransferHandler{store: store}
}

func (h *ProcessTransferHandler) Handle(req map[string]string, res ProcessTransferPresenter) {
	originID := req["originId"]
	destinationID := req["destinationId"]

	if originID == destinationID {
		res.TransferToSameAccount(originID)
		return
	}

	amount, err := strconv.ParseInt(req["amount"], 10, 64)

	if err != nil {
		res.InvalidParam("amount", req["amount"], "must be a number")
		return
	}

	if originID == "" {
		res.InvalidParam("originId", originID, "must not be empty")
		return
	}

	if destinationID == "" {
		res.InvalidParam("destinationId", destinationID, "must not be empty")
		return
	}

	if amount <= 0 {
		res.InvalidParam("amount", amount, "must be greater than zero")
		return
	}

	fromAccount, ok := h.store.Get(originID)

	if !ok {
		res.AccountNotFound(originID)
		return
	}

	toAccount, ok := h.store.Get(destinationID)

	if !ok {
		res.TransferToNonExistingAccount(destinationID)
		return
	}

	if fromAccount["balance"].(int64) < amount {
		res.InsufficientFunds(originID)
		return
	}

	fromAccount["balance"] = fromAccount["balance"].(int64) - amount
	fromAccount["transactions"] = append(fromAccount["transactions"].([]map[string]any), map[string]any{"amount": -amount, "type": "transfer", "to": destinationID})
	h.store.Put(originID, fromAccount)

	toAccount["balance"] = toAccount["balance"].(int64) + amount
	toAccount["transactions"] = append(toAccount["transactions"].([]map[string]any), map[string]any{"amount": amount, "type": "transfer", "from": originID})
	h.store.Put(destinationID, toAccount)

	res.TransferProcessed(fromAccount, toAccount, amount)
}
