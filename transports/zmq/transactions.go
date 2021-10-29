package zmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/libsv/nfwd"
	"github.com/ordishs/go-bitcoin"
	"github.com/pkg/errors"
)

type transactionHandler struct {
	svc    nfwd.TransactionService
	bc     chan []string
	bcdone chan bool
}

func NewHeadersHandler(svc nfwd.TransactionService) *transactionHandler {
	return &transactionHandler{
		svc: svc,
		bc:  make(chan []string),
	}
}

// Register will setup zmq with a handler
func (h *transactionHandler) Register(z *bitcoin.ZMQ) {
	if err := z.Subscribe("rawtx", h.bc); err != nil {
		log.Fatalln(err)
	}
}

func (h *transactionHandler) Listen() {
	fmt.Println("listening")
	for {
		select {
		case rawTx := <-h.bc:
			go func() {
				fmt.Println("handling message")
				ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
				defer cancelFn()
				fmt.Println(fmt.Sprintf("%+V", rawTx))
				if err := h.svc.Create(ctx, nfwd.TransactionCreate{TxHex: rawTx[1]}); err != nil {
					log.Println(err)
					return
				}
			}()
		case <-h.bcdone:
			close(h.bc)
		}
	}
}

func (h *transactionHandler) Close(z *bitcoin.ZMQ) error {
	if err := z.Unsubscribe("rawtx", h.bc); err != nil {
		return errors.WithMessage(err, "failed to unsubscribe rawtx")
	}
	close(h.bc)
	return nil
}
