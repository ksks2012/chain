package blocker

import (
	"log"

	"github.com/block-chain/global"
)

func GenGenesisBlock(previousHash []byte, difficulty int, miner string, minerRewards int) {
	log.Printf("Create genesis block...")
	newBlock := global.Block{
		PreviousHash: previousHash,
		Difficulty:   difficulty,
		Miner:        miner,
		MinerRewards: minerRewards,
	}
	newBlock.Hash = global.GetHash(newBlock, 0)
	global.MainChain.New(newBlock)
}
