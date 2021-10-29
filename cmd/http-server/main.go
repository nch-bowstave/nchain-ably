package main

import (
	"os"
	"os/signal"

	"github.com/ably/ably-go/ably"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/libsv/nfwd/config"
	"github.com/libsv/nfwd/config/zmq"
	txAbly "github.com/libsv/nfwd/data/ably"
	"github.com/libsv/nfwd/service"
	zmqTransport "github.com/libsv/nfwd/transports/zmq"
)

const appname = "node-forwarder"

func main() {
	log.Info().Msgf("starting %s\n", appname)
	config.SetDefaults()
	cfg := config.NewViperConfig(appname).
		WithServer().
		WithDb().
		WithDeployment(appname).
		WithLog().
		WithBitcoinNode().
		WithAbly().
		WithHeaderClient()

	if err := cfg.Validate(); err != nil {
		log.Fatal().Msgf("%s", err)
	}
	lvl, err := zerolog.ParseLevel(cfg.Logging.Level)
	if err != nil {
		log.Err(err).Msg("failed to parse")
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(lvl)
	}
	log.Info().Msg("connecting to ably")
	client, err := ably.NewRealtime(
		ably.WithKey(cfg.Ably.Key),
		ably.WithClientID(cfg.Ably.Username))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to ably")
	}
	log.Info().Msg("connected to ably")
	log.Info().Msg("connecting to node")
	zmqSub := zmq.Setup(cfg.Node)
	log.Info().Msg("node connected")

	txStore := txAbly.NewAbly(*cfg.Ably, client)
	txService := service.NewTransactionService(txStore)
	zmqHandler := zmqTransport.NewHeadersHandler(txService)
	zmqHandler.Register(zmqSub)
	log.Info().Msg("starting zmpq listener")

	go zmqHandler.Listen()
	log.Info().Msg("started zmpq listener")
	defer zmqHandler.Close(zmqSub)

	// wait for shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Info().Msg("shutting down node-listener")
}
