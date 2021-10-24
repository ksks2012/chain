package global

import "time"

type Block struct {
	Previous_hash string
	Hash          string
	Difficulty    string
	Nonce         int
	Timestamp     time.Time
	Transactions  []Transaction
	Miner         string
	Miner_rewards string
}
