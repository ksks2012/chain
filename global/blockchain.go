package global

type BlockChain struct {
	AdjustDifficultyBlocks int
	Difficulty             int
	BlockTime              int
	MiningRewards          int
	BlockLimitation        int
	Chain                  []Block
	PendingTransactions    []Transaction
}
