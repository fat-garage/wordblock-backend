package gipfs

import (
	"github.com/fat-garage/wordblock-backend/pkg/conf"
	"github.com/web3-storage/go-w3s-client"
	"log"
)

var W3sClient w3s.Client

func InitWeb3StorageClient() {
	cfg := conf.Cfg.Web3Storage

	c, err := w3s.NewClient(w3s.WithToken(cfg.Token))
	if err != nil {
		log.Fatalln(err)
	}
	W3sClient = c
}
