package global

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	PreviousHash []byte
	Hash         []byte
	Difficulty   int
	Nonce        int
	Timestamp    int64
	Transactions []Transaction
	Miner        string
	MinerRewards int
}

func (b *Block) New(previousHash []byte, difficulty int, miner string, minerRewards int) {
	b.PreviousHash = previousHash
	b.Difficulty = difficulty
	b.Nonce = 0
	b.Timestamp = time.Now().Unix()
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
	hashInput := fmt.Sprintf("%v%v%v%v", block.PreviousHash, block.Timestamp, getTransactionsString(block), strconv.Itoa(nonce))
	sha1Hasher.Write([]byte(hashInput))
	bs = sha1Hasher.Sum(nil)
	return

}
