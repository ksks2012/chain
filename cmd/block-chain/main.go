package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/block-chain/global"
	"github.com/block-chain/pkg/blocker"
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
	difficulty, err := strconv.Atoi(Diff)
	if err != nil {
		difficulty = 1
	}
	newBlock := blocker.GenGenesisBlock([]byte("Hello Chain!"), difficulty, "hong", 1)
	global.MainChain.New(newBlock)
	for i := 0; i <= 10; i++ {
		global.MainChain.MineBlock("hong")
	}

}

func setupFlag() error {
	flag.StringVar(&cfg, "config", "etc/", "指定要使用的設定檔路徑")
	flag.StringVar(&Diff, "Diff", "", "初始難度")
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
