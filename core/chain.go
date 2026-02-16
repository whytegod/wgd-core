package core

import "fmt"

type Blockchain struct {
	Blocks      []Block
	Mempool     []Transaction
	BlockReward float64
}

func NewBlockchain() *Blockchain {
	genesis := Block{
		Index:        0,
		Transactions: []Transaction{},
		PrevHash:     "",
		Hash:         "genesis_hash",
	}

	return &Blockchain{
		Blocks:      []Block{genesis},
		Mempool:     []Transaction{},
		BlockReward: 50,
	}
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	if tx.Verify() {
		bc.Mempool = append(bc.Mempool, tx)
	} else {
		fmt.Println("Invalid transaction")
	}
}

func (bc *Blockchain) MineBlock(miner string) {
	rewardTx := Transaction{
		From:   "network",
		To:     miner,
		Amount: bc.BlockReward,
	}

	bc.Mempool = append(bc.Mempool, rewardTx)

	newBlock := Block{
		Index:        len(bc.Blocks),
		Transactions: bc.Mempool,
		PrevHash:     bc.Blocks[len(bc.Blocks)-1].Hash,
		Hash:         fmt.Sprintf("hash_%d", len(bc.Blocks)),
	}

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Mempool = []Transaction{}

	fmt.Println("Block mined successfully")
}