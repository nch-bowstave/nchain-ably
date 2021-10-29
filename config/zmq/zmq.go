package zmq

import (
	"github.com/libsv/nfwd/config"
	"github.com/ordishs/go-bitcoin"
)

// Setup will setup the bitcoin node and zmq connections.
func Setup(c *config.BitcoinNode) *bitcoin.ZMQ {
	return bitcoin.NewZMQWithRaw(c.Host,  c.Port)

}
