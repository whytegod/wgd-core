package ledger

import (
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	Reward    float64
}

type Ledger struct {
	totalSupply float64
	balances    map[string]float64
	blocks      []Block
	height      int
	blockReward float64
}

/*
   Create New Blockchain Ledger
*/
func NewLedger(initialSupply float64, reward float64) *Ledger {

	l := &Ledger{
		totalSupply: initialSupply,
		balances:    make(map[string]float64),
		blocks:      []Block{},
		height:      0,
		blockReward: reward,
	}

	// Assign initial supply to treasury
	l.balances["treasury"] = initialSupply

	// Create Genesis Block
	genesis := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		Reward:    0,
	}

	l.blocks = append(l.blocks, genesis)
	l.height = 1

	return l
}

/*
   Mine New Block
*/
func (l *Ledger) AddBlock(data string) {

	newBlock := Block{
		Index:     l.height,
		Timestamp: time.Now(),
		Data:      data,
		Reward:    l.blockReward,
	}

	l.blocks = append(l.blocks, newBlock)
	l.height++

	// Mint reward to treasury
	l.totalSupply += l.blockReward
	l.balances["treasury"] += l.blockReward
}

/*
   Getters
*/

func (l *Ledger) TotalSupply() float64 {
	return l.totalSupply
}

func (l *Ledger) Height() int {
	return l.height
}

func (l *Ledger) TreasuryBalance() float64 {
	return l.balances["treasury"]
}

func (l *Ledger) PrintChain() {
	for _, block := range l.blocks {
		fmt.Println("-------------")
		fmt.Println("Block:", block.Index)
		fmt.Println("Time:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("Reward:", block.Reward)
	}
}