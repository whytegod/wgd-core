package ledger

import "fmt"

type Ledger struct {
	totalSupply float64
	balances    map[string]float64
}

func NewLedger(initialSupply float64) *Ledger {
	l := &Ledger{
		totalSupply: initialSupply,
		balances:    make(map[string]float64),
	}

	// Assign all genesis supply to treasury
	l.balances["treasury"] = initialSupply

	return l
}

func (l *Ledger) BalanceOf(account string) float64 {
	return l.balances[account]
}

func (l *Ledger) Transfer(from, to string, amount float64) error {
	if l.balances[from] < amount {
		return fmt.Errorf("insufficient balance")
	}

	l.balances[from] -= amount
	l.balances[to] += amount

	return nil
}

func (l *Ledger) TotalSupply() float64 {
	return l.totalSupply
}