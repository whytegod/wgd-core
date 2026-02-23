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
	mu    sync.RWMutex
	store map[UTXOKey]tx.TxOutput
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		store: make(map[UTXOKey]tx.TxOutput),
	}
}

func (u *UTXOSet) Get(key UTXOKey) (tx.TxOutput, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	val, ok := u.store[key]
	return val, ok
}

func (u *UTXOSet) Put(key UTXOKey, out tx.TxOutput) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.store[key] = out
}

func (u *UTXOSet) Delete(key UTXOKey) {
	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.store, key)
}