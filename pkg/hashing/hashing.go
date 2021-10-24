package hashing

import (
	"crypto/sha1"

	"github.com/block-chain/global"
)

var aaa int

func transactionToString(transaction global.Transaction) (transactionstring string) {
	transactionstring = transaction.Sender + transaction.Receiver + transaction.Amounts + transaction.Fee + transaction.Message
	return
}

func getTransactionsString(block global.Block) (transactionsString string) {
	transactionsString = ""
	for _, transaction := range block.Transactions {
		transactionsString += transactionToString(transaction)
	}
	return
}

func GetHash(block global.Block, nonce string) (bs []byte) {
	sha1Hasher := sha1.New()
	hashInput := block.Previous_hash + block.Timestamp.String() + getTransactionsString(block) + nonce
	sha1Hasher.Write([]byte(hashInput))
	bs = sha1Hasher.Sum(nil)
	return

}
