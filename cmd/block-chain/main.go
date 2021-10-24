package main

import (
	"flag"
	"log"
	"strings"

	"github.com/block-chain/global"
	"github.com/block-chain/pkg/hashing"
	"github.com/block-chain/pkg/setting"
)

var (
	cfg string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	next_hash := hashing.GetHash(global.Block{}, "123")
	log.Printf("next_hash: %s", next_hash)
}

func setupFlag() error {
	flag.StringVar(&cfg, "config", "etc/", "指定要使用的設定檔路徑")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(cfg, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("BlockChain", &global.BlockChainSetting)
	if err != nil {
		return err
	}
	return nil
}
