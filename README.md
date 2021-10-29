# ably-bitcoin-indexer

[![Release](https://img.shields.io/github/release-pre/nch-bowstave/nchain-ably.svg?logo=github&style=flat&v=1)](https://github.com/nch-bowstave/nchain-ably/releases)
[![Build Status](https://img.shields.io/github/workflow/status/nch-bowstave/nchain-ably/go?logo=github&v=3)](https://github.com/nch-bowstave/nchain-ably/actions)
[![Report](https://goreportcard.com/badge/github.com/nch-bowstave/nchain-ably?style=flat&v=1)](https://goreportcard.com/report/github.com/nch-bowstave/nchain-ably)
[![Go](https://img.shields.io/github/go-mod/go-version/nch-bowstave/nchain-ably?v=1)](https://golang.org/)

This is a service that can be used as a new data source for Ably, a real time messaging platform.

This is built using GoLang and will listen to a node and forward all transactions to Ably.

As well as transactions, Bitcoin SV also defines several data protocols that can be sued by applications to 
standardise encoding etc. We handle the following data protocols:

- bitcom
- bitcom-b
- bitcom-d
- run
- metanet

## Quick Start

You will need a bitcoin sv node to get started, contact us for info on setting one up.

Using docker is the quickest way to get started.

The only settings you need to be concerned with are listed in the [docker-compose](docker-compose.yml) file:

      - ABLY_KEY=yourkey
      - NODE_HOST=localhost
      - NODE_USER=youruser
      - NODE_PASSWORD=yoursecret

You will need to build a local image before running:

```bash
make build-image
```

Then once built, do:

```bash
make run-compose
```

### Go run

You can also run with go run:

```shell
 ABLY_KEY=<YOUR KEY> NODE_HOST=<YOUR NODE> NODE_USER=<YOUR USER> NODE_PASSWORD=<YOUR PASSWORD> go run cmd/http-server/main.go
```

## Channels

We currently send data to several channels:

- txraw - raw transaction data as they come in

The below channels contain only data from outputs from transactions that match these protocols:

- bitcom
- bitcom-b
- bitcom-d
- run
- metanet

The data format is:

```shell
{
    "txId":"transactionID",
    "script":"protocol script hex"
    "protocol":"bitcom/run/etc"
}
```
