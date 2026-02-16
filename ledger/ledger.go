cat > ledger/ledger.go << 'EOF'
package ledger

type Ledger struct {
	totalSupply     uint64
	treasuryBalance uint64
}

func NewLedger(initialSupply uint64) *Ledger {
	return &Ledger{
		totalSupply:     initialSupply,
		treasuryBalance: initialSupply,
	}
}

func (l *Ledger) TotalSupply() uint64 {
	return l.totalSupply
}

func (l *Ledger) TreasuryBalance() uint64 {
	return l.treasuryBalance
}
EOF