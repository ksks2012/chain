package main

import (
	"flag"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/block-chain/global"
	"github.com/block-chain/pkg/blocker"
	"github.com/block-chain/pkg/setting"
)

var (
	cfg   string
	Diff  string
	miner string
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
	blocker.GenAddress()
	difficulty, err := strconv.Atoi(Diff)
	if err != nil {
		difficulty = 1
	}
	var newBlock global.Block
	newBlock.GenGenesisBlock([]byte("Hello Chain!"), difficulty, "hong", 10)
	global.MainChain.New(newBlock)
	for i := 0; i <= 3; i++ {
		// Step1: initialize a transaction
		transaction := global.InitialTransaction(
			// string(rsakey.RsaSignWithSha256([]byte(miner), blocker.PrivateKey)),
			blocker.PublicKey,
			"test123",
			1,
			1,
			"Test")
		if !reflect.DeepEqual(transaction, global.Transaction{}) {
			// Step2: Sign your transaction
			signature := global.SignTransaction(transaction)
			// Step3: Send it to blockchain
			global.MainChain.AddTransaction(transaction, signature)
			global.MainChain.MineBlock(miner)
			global.MainChain.AdjustDifficulty()
			log.Printf("Surplus %v", global.MainChain.GetSurplus(blocker.PublicKey))
		}
	}

}

func setupFlag() error {
	flag.StringVar(&cfg, "config", "etc/", "指定要使用的設定檔路徑")
	flag.StringVar(&Diff, "Diff", "", "初始難度")
	flag.StringVar(&miner, "miner", "xxx", "挖掘者")
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
