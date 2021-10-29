package nfwd

import (
	"context"

	"github.com/libsv/go-bt/v2"
)

// TransactionCreate contains the properties required
// to create a transaction.
type TransactionCreate struct {
	TxHex string
}

// TransactionService enforces validation of arguments and business rules.
type TransactionService interface {
	// Create will store a transaction in the db.
	Create(ctx context.Context, req TransactionCreate) error
}

// TransactionWriter will add transactions to a data store.
type TransactionWriter interface {
	// Create will add a transaction to the data store.
	Create(ctx context.Context, req *bt.Tx) error
}
