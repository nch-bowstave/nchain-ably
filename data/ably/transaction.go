package ably

import (
	"context"
	"encoding/hex"

	"github.com/ably/ably-go/ably"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/libsv/nfwd/config"
)

var protocols = map[string]string{
	"3137335a6659393779374e6a625a326b413373796a53744363444e41786276564438": "bitcom",
	"31394878696756345179427633744870515663554551797131707a5a56646f417574": "bitcom-b",
	"3139694733575459537362796f7333754a373333794b347a45696f69314665734e55": "bitcom-d",
	"72756e":   "run",
	"6d657461": "metanet",
}

type Output struct {
	TxID     string `json:"txId"`
	Script   string `json:"script"`
	Protocol string `json:"protocol"`
}

type ablyTx struct {
	cfg    config.Ably
	client *ably.Realtime
	chs    map[string]*ably.RealtimeChannel
}

func NewAbly(cfg config.Ably, client *ably.Realtime) *ablyTx {
	a := &ablyTx{
		cfg:    cfg,
		client: client,
		chs: map[string]*ably.RealtimeChannel{
			"txfeed": client.Channels.Get("rawtx"),
		},
	}
	for _, v := range protocols {
		a.chs[v] = client.Channels.Get(v)
	}
	return a
}

// Create will post all transactions to a rawtx message type on the txfeed channel.
//
// It will check all outputs for known protocols and then if a known protocol
// send the output hex to the txfeed channel with the message type of protocol.
func (a *ablyTx) Create(ctx context.Context, req *bt.Tx) error {
	log.Debug().Msg("received new ably tx, broadcasting to rawtx channel...")
	if int64(len(req.String())) < a.cfg.MaxMessage {
		if err := a.chs["txfeed"].Publish(ctx, "rawtx", req.String()); err != nil {
			return errors.Wrap(err, "failed to publish rawtx to ably")
		}
		log.Debug().Msg("broadcast to rawtx channel success")
	} else {
		log.Debug().Msgf("tx exceeds max allowed size for ably '%d' bytes, skipping", len(req.String()))
	}

	for _, o := range req.Outputs {
		if !o.LockingScript.IsData() || int64(len(o.LockingScript.String())) >= a.cfg.MaxMessage {
			log.Debug().Msgf("output exceeds max allowed size for ably '%d' bytes, skipping", len(*o.LockingScript))
			continue
		}
		log.Debug().Msg("found data output, parsing for known protocol")
		pp, err := bscript.DecodeParts(*o.LockingScript)
		if err != nil {
			return err
		}
		if len(pp) <= 3 {
			continue
		}
		proto, ok := protocols[hex.EncodeToString(pp[2])]
		if !ok {
			continue
		}
		log.Debug().Msgf("broadcasting to '%s' channel", proto)
		if err := a.chs[proto].Publish(ctx, proto, Output{
			TxID:     req.TxID(),
			Script:   o.LockingScriptHexString(),
			Protocol: proto,
		}); err != nil {
			return errors.Wrapf(err, "failed to publish '%s' message to ably", proto)
		}
		log.Debug().Msgf("broadcast to '%s' channel success", proto)
	}
	return nil
}
