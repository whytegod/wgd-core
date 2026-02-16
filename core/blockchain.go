package core

type Blockchain struct {
	Blocks     []Block
	Difficulty int
	Mempool    []Transaction
}

func NewBlockchain() *Blockchain {
	genesis := Block{
		Index:     0,
		Timestamp: "Genesis",
		Hash:      "0000genesis",
	}

	return &Blockchain{
		Blocks:     []Block{genesis},
		Difficulty: 3,
		Mempool:    []Transaction{},
	}
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	bc.Mempool = append(bc.Mempool, tx)
}

func (bc *Blockchain) MinePendingTransactions() {
	prev := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(
		prev.Index+1,
		bc.Mempool,
		prev.Hash,
		bc.Difficulty,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Mempool = []Transaction{}
}