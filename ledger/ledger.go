package ledger

import "fmt"

type Block struct {
	Index  int
	Data   string
	Reward float64
}

type Ledger struct {
	totalSupply float64
	balances    map[string]float64
	blocks      []Block
	height      int
}

func NewLedger(initialSupply float64) *Ledger {

	l := &Ledger{
		totalSupply: initialSupply,
		balances:    make(map[string]float64),
		blocks:      []Block{},
		height:      0,
	}

	// Assign genesis supply to treasury
	l.balances["treasury"] = initialSupply

	// Create Genesis Block
	genesis := Block{
		Index:  0,
		Data:   "Genesis Block",
		Reward: 0,
	}

	l.blocks = append(l.blocks, genesis)

	return l
}