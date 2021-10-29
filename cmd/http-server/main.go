package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/libsv/nfwd/config"
	"github.com/libsv/nfwd/config/zmq"
	"github.com/libsv/nfwd/data/elastic"
	"github.com/libsv/nfwd/service"
	zmqTransport "github.com/libsv/nfwd/transports/zmq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/elastic/go-elasticsearch/v7"
)

const appname = "node-forwarder"

func main() {
	log.Printf("starting %s\n", appname)
	config.SetDefaults()
	cfg := config.NewViperConfig(appname).
		WithServer().
		WithDb().
		WithDeployment(appname).
		WithLog().
		WithBitcoinNode().
		WithWoc().
		WithHeaderClient()

	if err := cfg.Validate(); err != nil {
		log.Fatalf("%s", err)
	}
	lvl, err := zerolog.ParseLevel(cfg.Logging.Level)
	if err != nil {
		log.Println(err)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(lvl)
	}
	// TODO elastic setup
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://elastic:9200",
		},
	})

	fmt.Println("connecting to node")
	zmqSub := zmq.Setup(cfg.Node)
	fmt.Println("node connected")

	txStore := elastic.NewTransactions(es)
	txService := service.NewTransactionService(txStore)
	zmqHandler := zmqTransport.NewHeadersHandler(txService)
	zmqHandler.Register(zmqSub)
	fmt.Println("starting zmpq listener")

	go zmqHandler.Listen()
	fmt.Println("started zmpq listener")
	defer zmqHandler.Close(zmqSub)

	// wait for shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<- c
	fmt.Println("shutting down node-listener")
}
