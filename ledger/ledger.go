package ledger

// Ledger implements monetary policy bookkeeping for WGD
type Ledger struct {
	totalSupply     uint64 // expressed in Werts (smallest unit)
	treasuryBalance uint64 // treasury disabled: will remain 0
}

const (
	// supply and unit definitions
	WGDTotal      uint64 = 9720000               // 9,720,000 WGD
	WertsPerWGD   uint64 = 100_000_000            // 1 WGD = 100,000,000 Werts (like satoshi)
	TotalWerts    uint64 = WGDTotal * WertsPerWGD // total supply in Werts
)

// NewLedger returns ledger with immutable total supply and zero treasury.
func NewLedger() *Ledger {
	return &Ledger{
		totalSupply:     TotalWerts,
		treasuryBalance: 0,
	}
}

// TotalSupply returns supply in Werts.
func (l *Ledger) TotalSupply() uint64 {
	return l.totalSupply
}

// TreasuryBalance returns treasury balance in Werts (should be zero).
func (l *Ledger) TreasuryBalance() uint64 {
	return l.treasuryBalance
}