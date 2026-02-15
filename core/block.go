package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

func (b *Block) CalculateHash() string {
	record := fmt.Sprintf("%d%s%v%s",
		b.Index,
		b.Timestamp,
		b.Transactions,
		b.PrevHash,
	)

	h := sha256.Sum256([]byte(record))
	return hex.EncodeToString(h[:])
}

func NewBlock(index int, txs []Transaction, prevHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: txs,
		PrevHash:     prevHash,
	}

	block.Hash = block.CalculateHash()
	return block
}