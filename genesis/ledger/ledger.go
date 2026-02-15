package ledger

type Account struct {
	Address string
	Balance float64
}

type Ledger struct {
	Accounts map[string]*Account
}

func NewLedger() *Ledger {
	return &Ledger{
		Accounts: make(map[string]*Account),
	}
}

func (l *Ledger) CreateAccount(address string, balance float64) {
	l.Accounts[address] = &Account{
		Address: address,
		Balance: balance,
	}
}

func (l *Ledger) GetBalance(address string) float64 {
	if acc, exists := l.Accounts[address]; exists {
		return acc.Balance
	}
	return 0
}