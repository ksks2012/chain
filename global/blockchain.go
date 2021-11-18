package global

import (
	"bytes"
	"log"
	"sort"
	"time"
)

type BlockChain struct {
	AdjustDifficultyBlocks int
	Difficulty             int
	BlockTime              int
	MiningRewards          int
	BlockLimitation        int
	Chain                  []Block
	PendingTransactions    []Transaction
}

var (
	MainChain BlockChain
)

func (bc *BlockChain) AddTransactionToBlock(block Block) {
	//  Get the transaction with highest fee by block_limitation
	sort.SliceStable(bc.PendingTransactions, func(i, j int) bool {
		return bc.PendingTransactions[i].Fee < bc.PendingTransactions[j].Fee
	})
	transcationAccepted := []Transaction{}
	if len(bc.PendingTransactions) > bc.BlockLimitation {
		transcationAccepted = bc.PendingTransactions[:bc.BlockLimitation]
		bc.PendingTransactions = bc.PendingTransactions[bc.BlockLimitation:]
	} else {
		transcationAccepted = bc.PendingTransactions
		bc.PendingTransactions = []Transaction{}
	}
	block.Transactions = transcationAccepted
}

func (bc *BlockChain) MineBlock(miner string) {
	// mine new block and add to chain
	log.Printf("MineBlock")
	startTime := time.Now().Unix()

	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{}
	newBlock.New(lastBlock.Hash, bc.Difficulty, miner, bc.MiningRewards)

	bc.AddTransactionToBlock(newBlock)
	newBlock.PreviousHash = lastBlock.Hash
	newBlock.Difficulty = bc.Difficulty
	newBlock.Hash = GetHash(newBlock, newBlock.Nonce)

	nonce := make([]byte, bc.Difficulty)
	for ; !bytes.Equal(newBlock.Hash[0:bc.Difficulty], nonce); newBlock.Nonce++ {
		newBlock.Hash = GetHash(newBlock, newBlock.Nonce)
	}

	timeConsumed := time.Now().Unix() - startTime
	log.Printf("Hash found: %x @ diffuculty %v, time cost: %v", []byte(newBlock.Hash), bc.Difficulty, timeConsumed)
	bc.Chain = append(bc.Chain, newBlock)
}
