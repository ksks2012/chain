package service

import (
	"crypto/sha1"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/block-chain/pkg/blocker"
	"github.com/block-chain/pkg/rsakey"
)

type Block struct {
	PreviousHash []byte
	Hash         []byte
	Difficulty   int
	Nonce        int
	Timestamp    int64
	Transactions []Transaction
	Miner        []byte
	MinerRewards int64
}

func (b *Block) New(previousHash []byte, difficulty int, miner string, minerRewards int64) {
	b.PreviousHash = previousHash
	b.Difficulty = difficulty
	b.Nonce = 0
	b.Timestamp = time.Now().Unix()
	b.Transactions = []Transaction{}
	b.Miner = blocker.PublicKey
	b.MinerRewards = minerRewards
}

func (b *Block) GenGenesisBlock(previousHash []byte, difficulty int, miner string, minerRewards int64) {
	log.Printf("Create genesis block...")
	b.New(previousHash, difficulty, miner, minerRewards)
	b.Hash = GetHash(*b, 0)
}

func getTransactionsString(block Block) (transactionsString string) {
	transactionsString = ""
	for _, transaction := range block.Transactions {
		transactionsString += transaction.transactionToString()
	}
	return
}

func GetHash(block Block, nonce int) (bs []byte) {
	sha1Hasher := sha1.New()
	hashInput := fmt.Sprintf("%v%v%v%v", block.PreviousHash, block.Timestamp, getTransactionsString(block), strconv.Itoa(nonce))
	sha1Hasher.Write([]byte(hashInput))
	bs = sha1Hasher.Sum(nil)
	return

}

func SignTransaction(transaction Transaction) []byte {
	transactionString := transaction.transactionToString()
	return rsakey.RsaSignWithSha256([]byte(transactionString), blocker.PrivateKey)
}
