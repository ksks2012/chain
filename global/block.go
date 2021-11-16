package global

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	PreviousHash string
	Hash         string
	Difficulty   int
	Nonce        int
	Timestamp    time.Time
	Transactions []Transaction
	Miner        string
	MinerRewards int
}

func (b *Block) New(previousHash string, difficulty int, miner string, minerRewards int) {
	b.PreviousHash = previousHash
	b.Difficulty = difficulty
	b.Nonce = 0
	b.Timestamp = time.Now()
	b.Transactions = []Transaction{}
	b.Miner = miner
	b.MinerRewards = minerRewards
}

func transactionToString(transaction Transaction) (transactionstring string) {
	transactionstring = transaction.Sender + transaction.Receiver + transaction.Amounts + transaction.Fee + transaction.Message
	return
}

func getTransactionsString(block Block) (transactionsString string) {
	transactionsString = ""
	for _, transaction := range block.Transactions {
		transactionsString += transactionToString(transaction)
	}
	return
}

func GetHash(block Block, nonce int) (bs []byte) {
	sha1Hasher := sha1.New()
	hashInput := fmt.Sprintf("%v%v%v%v", block.PreviousHash, block.Timestamp.Unix(), getTransactionsString(block), strconv.Itoa(nonce))
	sha1Hasher.Write([]byte(hashInput))
	bs = sha1Hasher.Sum(nil)
	return

}
