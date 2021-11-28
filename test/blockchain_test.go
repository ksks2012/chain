package test

import (
	"log"
	"testing"
)

func TestVerifyBlockChain(t *testing.T) {
	var blockChain service.BlockChain
	var newBlock service.Block
	newBlock.GenGenesisBlock([]byte(""), 1, "hong", 1)
	blockChain.New(newBlock)
	log.Printf("%v %x", len(blockChain.Chain), blockChain.Chain[0].Hash)
	blockChain.MineBlock("hong")
	blockChain.VerifyBlockchain()
	log.Print("Insert fake transaction.")
	fakeTransaction := service.Transaction{[]byte("test123"), "test456", 100, 1, "TEST"}
	blockChain.Chain[0].Transactions = append(blockChain.Chain[0].Transactions, fakeTransaction)
	blockChain.MineBlock("hong")
	blockChain.VerifyBlockchain()
}
