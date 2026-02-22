package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
	"strconv"
)

// Block represents a single block in the chain.
type Block struct {
	Index        int
	Timestamp    int64
	Transactions []*Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int
	Difficulty   int
}

// NewBlock creates and returns a mined block pointer.
func NewBlock(index int, txs []*Transaction, prevHash []byte, difficulty int) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
		Difficulty:   difficulty,
	}

	block.Mine()
	return block
}

// calculateHash produces a SHA256 hash for the block content.
func (b *Block) calculateHash() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)

	// encode transactions IDs only (avoid including signatures for deterministic hash)
	txIDs := make([][]byte, 0, len(b.Transactions))
	for _, tx := range b.Transactions {
		if tx == nil {
			txIDs = append(txIDs, []byte{})
		} else {
			txIDs = append(txIDs, tx.ID)
		}
	}

	_ = enc.Encode(b.Index)
	_ = enc.Encode(b.Timestamp)
	_ = enc.Encode(txIDs)
	_ = enc.Encode(b.PrevHash)
	_ = enc.Encode(b.Nonce)
	_ = enc.Encode(b.Difficulty)

	sum := sha256.Sum256(encoded.Bytes())
	return sum[:]
}

// Mine searches for a hash with a simple difficulty target (leading zeros in hex)
func (b *Block) Mine() {
	targetPrefix := ""
	// difficulty expresses number of leading hex '0' nibbles required.
	for i := 0; i < b.Difficulty; i++ {
		targetPrefix += "0"
	}

	for {
		hash := b.calculateHash()
		hex := fmt.Sprintf("%x", hash)
		if len(targetPrefix) == 0 || (len(hex) >= len(targetPrefix) && hex[:len(targetPrefix)] == targetPrefix) {
			b.Hash = hash
			return
		}
		b.Nonce++
	}
}