package global

import (
	"bytes"
	"log"
	"math"
	"sort"
	"time"

	"github.com/block-chain/pkg/blocker"
	"github.com/block-chain/pkg/rsakey"
)

type BlockChain struct {
	AdjustDifficultyBlocks int
	Difficulty             int
	BlockTime              int
	MiningRewards          int64
	BlockLimitation        int
	Chain                  []Block
	PendingTransactions    []Transaction
}

var (
	MainChain BlockChain
)

func (bc *BlockChain) New(initBlock Block) {
	bc.AdjustDifficultyBlocks = 5
	bc.Difficulty = initBlock.Difficulty
	bc.MiningRewards = initBlock.MinerRewards
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
	for !bytes.Equal(newBlock.Hash[0:bc.Difficulty], nonce) {
		newBlock.Nonce++
		newBlock.Hash = GetHash(newBlock, newBlock.Nonce)
	}
	// log.Printf("nonce %v %v", newBlock.Hash[0:bc.Difficulty], newBlock.Hash[0:(bc.Difficulty*2)])
	newBlock.Timestamp = time.Now().Unix()
	timeConsumed := time.Now().Unix() - startTime
	log.Printf("Hash found: %x @ diffuculty %v, time cost: %vs", []byte(newBlock.Hash), bc.Difficulty, timeConsumed)
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *BlockChain) AdjustDifficulty() int {
	chainLength := len(bc.Chain)
	if (chainLength%bc.AdjustDifficultyBlocks != 1) && chainLength <= bc.AdjustDifficultyBlocks {
		return bc.Difficulty
	}

	start := bc.Chain[chainLength-bc.AdjustDifficultyBlocks-1].Timestamp
	finish := bc.Chain[chainLength-1].Timestamp
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

func (bc *BlockChain) GetSurplus(account []byte) (surplus int64) {
	surplus = 0
	for _, block := range bc.Chain {
		miner := false
		if bytes.Equal(block.Miner, account) {
			miner = true
			surplus += block.MinerRewards
		}
		for _, transaction := range block.Transactions {
			if miner {
				surplus += transaction.Fee
			}
			if bytes.Equal([]byte(transaction.Sender), account) {
				surplus -= transaction.Amounts
				surplus -= transaction.Fee
			} else if bytes.Equal([]byte(transaction.Receiver), account) {
				surplus += transaction.Amounts
			}
		}
	}
	return surplus
}

func (bc *BlockChain) VerifyBlockchain() bool {
	previousHash := make([]byte, 0)
	for idx, block := range bc.Chain {
		if !bytes.Equal(GetHash(block, block.Nonce), block.Hash) {
			log.Printf("Error: Hash not matched!")
			return false
		} else if !bytes.Equal(previousHash, block.PreviousHash) && idx != 0 {
			log.Printf("Error: Hash not matched to previous hash!")
			return false
		}
		previousHash = block.Hash
	}
	log.Printf("Hash correct!")
	return true

}

func (bc *BlockChain) AddTransaction(transaction Transaction, signature []byte) bool {
	if transaction.Fee+transaction.Amounts > bc.GetSurplus(transaction.Sender) {
		return false
	}
	transactionString := transaction.transactionToString()
	if rsakey.RsaVerySignWithSha256(
		[]byte(transactionString), signature, blocker.PublicKey) {
		log.Printf("Authorized successfully!")
		bc.PendingTransactions = append(bc.PendingTransactions, transaction)
		return true
	}
	log.Printf("RSA Verified wrong!")
	return false
}
