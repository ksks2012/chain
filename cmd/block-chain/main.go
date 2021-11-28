package main

import (
	"flag"
	"log"
	"strings"

	"github.com/block-chain/global"
	"github.com/block-chain/internal/service"
	"github.com/block-chain/pkg/setting"
)

var (
	cfg  string
	Diff string
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
	var newBlock service.Block
	newBlock.GenGenesisBlock([]byte("Hello Chain!"), global.BlockChainSetting.Difficulty, "hong", 1)
	service.MainChain.New(newBlock, *global.BlockChainSetting)
	for i := 0; i <= 10; i++ {
		service.MainChain.MineBlock("hong")
		service.MainChain.AdjustDifficulty()
	}

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
