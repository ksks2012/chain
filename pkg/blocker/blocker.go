package blocker

import (
	"log"

	"github.com/block-chain/global"
)

func GenGenesisBlock(previousHash string, difficulty int, miner string, minerRewards int) {
	log.Printf("Create genesis block...")
	newBlock := global.Block{
		PreviousHash: previousHash,
		Difficulty:   difficulty,
		Miner:        miner,
		MinerRewards: minerRewards,
	}
	newBlock.Hash = string(global.GetHash(newBlock, 0))
	global.MainChain.Difficulty = newBlock.Difficulty
	global.MainChain.Chain = append(global.MainChain.Chain, newBlock)
}
