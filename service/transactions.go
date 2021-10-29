package service

import (
	"context"

	"github.com/libsv/go-bt/v2"
	"github.com/libsv/nfwd"
	"github.com/pkg/errors"
	validator "github.com/theflyingcodr/govalidator"
)

type transactionService struct {
	txWtr nfwd.TransactionWriter
}

// NewTransactionService will setup and return a new transaction service, used to add and get transactions.
func NewTransactionService(txWtr nfwd.TransactionWriter) *transactionService {
	return &transactionService{
		txWtr: txWtr,
	}
}

// Create will validate a new transaction before adding to a data store.
func (h *transactionService) Create(ctx context.Context, req nfwd.TransactionCreate) error {
	tx, err := bt.NewTxFromString(req.TxHex)
	if err != nil {
		return validator.ErrValidation{
			"tx": []string{
				err.Error(),
			},
		}
	}
	if err := h.txWtr.Create(ctx, tx); err != nil {
		return errors.Wrapf(err, "failed to create transaction with id %s", tx.TxID())
	}
	return nil
}
