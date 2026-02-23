package ledger

import (
	"sync"
	"wgd-core/internal/tx"
)

type UTXOKey struct {
	TxHash [32]byte
	Index  uint32
}

type UTXOSet struct {
	mu sync.RWMutex
	set map[UTXOKey]tx.TxOutput
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		set: make(map[UTXOKey]tx.TxOutput),
	}
}

func (u *UTXOSet) Add(txHash [32]byte, index uint32, output tx.TxOutput) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.set[UTXOKey{txHash, index}] = output
}

func (u *UTXOSet) Spend(txHash [32]byte, index uint32) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	key := UTXOKey{txHash, index}
	if _, exists := u.set[key]; !exists {
		return false
	}
	delete(u.set, key)
	return true
}

func (u *UTXOSet) Get(txHash [32]byte, index uint32) (tx.TxOutput, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	out, ok := u.set[UTXOKey{txHash, index}]
	return out, ok
}