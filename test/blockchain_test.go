package test

import (
	"log"
	"testing"

	"github.com/block-chain/global"
	"github.com/block-chain/pkg/blocker"
)

func TestVerifyBlockChain(t *testing.T) {
	var blockChain global.BlockChain
	newBlock := blocker.GenGenesisBlock([]byte(""), 1, "hong", 1)
	blockChain.New(newBlock)
	log.Printf("%v %x", len(blockChain.Chain), blockChain.Chain[0].Hash)
	blockChain.MineBlock("hong")
	blockChain.VerifyBlockchain()
	log.Print("Insert fake transaction.")
	fakeTransaction := global.Transaction{"test123", "test456", 100, 1, "TEST"}
	blockChain.Chain[0].Transactions = append(blockChain.Chain[0].Transactions, fakeTransaction)
	blockChain.MineBlock("hong")
	blockChain.VerifyBlockchain()
}
