package core

import (
	"fmt"
)

type Blockchain struct {
	Blocks          []Block
	Difficulty      int
	BlockReward     float64
	HalvingInterval int
	Mempool         []Transaction
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
		Blocks:          []Block{genesis},
		Difficulty:      difficulty,
		BlockReward:     50,
		HalvingInterval: 10,
		Mempool:         []Transaction{},
	}
}

func (bc *Blockchain) calculateReward() float64 {
	halvings := len(bc.Blocks) / bc.HalvingInterval
	reward := bc.BlockReward

	for i := 0; i < halvings; i++ {
		reward /= 2
	}

	if reward < 0.0001 {
		reward = 0
	}

	return reward
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	if tx.Verify() {
		bc.Mempool = append(bc.Mempool, tx)
	}
}

func (bc *Blockchain) MineBlock(miner string) {

	rewardTx := Transaction{
		From:   "SYSTEM",
		To:     miner,
		Amount: bc.calculateReward(),
	}

	txs := append(bc.Mempool, rewardTx)

	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(
		prevBlock.Index+1,
		txs,
		prevBlock.Hash,
		bc.Difficulty,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Mempool = []Transaction{}

	fmt.Println("Block mined successfully")
}