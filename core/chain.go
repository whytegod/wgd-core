package core

import "fmt"

type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

func NewBlockchain() *Blockchain {
	difficulty := 3

	genesis := Block{
		Index:        0,
		Timestamp:    "GENESIS",
		Transactions: []Transaction{},
		PrevHash:     "0",
		Nonce:        0,
	}

	genesis.Hash = genesis.CalculateHash()

	return &Blockchain{
		Blocks:     []Block{genesis},
		Difficulty: difficulty,
	}
}

func (bc *Blockchain) AddBlock(txs []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(
		prevBlock.Index+1,
		txs,
		prevBlock.Hash,
		bc.Difficulty,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		if current.Hash != current.CalculateHash() {
			return false
		}

		if current.PrevHash != previous.Hash {
			return false
		}
	}

	return true
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Println("Block Index:", block.Index)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Nonce:", block.Nonce)
		fmt.Println("Tx Count:", len(block.Transactions))
		fmt.Println("--------------------------")
	}
}