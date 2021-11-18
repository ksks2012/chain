package global

import (
	"bytes"
	"log"
	"math"
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

func (bc *BlockChain) New(initBlock Block) {
	bc.AdjustDifficultyBlocks = 1
	bc.Difficulty = initBlock.Difficulty
	bc.Chain = append(bc.Chain, initBlock)
}

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
	// log.Printf("nonce %v", nonce)
	for ; !bytes.Equal(newBlock.Hash[0:bc.Difficulty], nonce); newBlock.Nonce++ {
		newBlock.Hash = GetHash(newBlock, newBlock.Nonce)
	}
	// log.Printf("nonce %v %v", newBlock.Hash[0:bc.Difficulty], newBlock.Hash[0:(bc.Difficulty*2)])
	newBlock.Timestamp = time.Now()
	timeConsumed := time.Now().Unix() - startTime
	log.Printf("Hash found: %x @ diffuculty %v, time cost: %vs", []byte(newBlock.Hash), bc.Difficulty, timeConsumed)
	bc.Chain = append(bc.Chain, newBlock)
	bc.adjustDifficulty()
}

func (bc *BlockChain) adjustDifficulty() int {
	chainLength := len(bc.Chain)
	if (chainLength%bc.AdjustDifficultyBlocks != 1) && chainLength <= bc.AdjustDifficultyBlocks {
		return bc.Difficulty
	}

	start := bc.Chain[chainLength-bc.AdjustDifficultyBlocks-1].Timestamp.Unix()
	finish := bc.Chain[chainLength-1].Timestamp.Unix()
	avgTimeConsumed := math.Round(float64(finish-start) / float64(bc.AdjustDifficultyBlocks))
	if avgTimeConsumed > float64(bc.BlockTime) {
		log.Printf("Average block time:%vs. Lower the difficulty", avgTimeConsumed)
		bc.Difficulty -= 1
	} else {
		log.Printf("Average block time:%vs. High up the difficulty", avgTimeConsumed)
		bc.Difficulty += 1
	}
	return bc.Difficulty
}
