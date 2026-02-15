package ledger

import "github.com/whytegod/wgd-core/block"

type Ledger struct {
	Blocks []block.Block
}

func NewLedger() *Ledger {
	return &Ledger{
		Blocks: []block.Block{},
	}
}

func (l *Ledger) AddBlock(b block.Block) {
	l.Blocks = append(l.Blocks, b)
}

func (l *Ledger) Height() int {
	return len(l.Blocks) - 1
}

func (l *Ledger) TotalSupply() float64 {
	total := 0.0
	for _, b := range l.Blocks {
		total += b.Reward
	}
	return total
}