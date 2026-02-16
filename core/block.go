package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) +
		b.Timestamp +
		b.PrevHash +
		strconv.Itoa(b.Nonce)

	for _, tx := range b.Transactions {
		record += tx.Hash
	}

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (b *Block) Mine(difficulty int) {
	target := strings.Repeat("0", difficulty)

	for {
		b.Hash = b.calculateHash()
		if b.Hash[:difficulty] == target {
			break
		}
		b.Nonce++
	}
}

func NewBlock(index int, txs []Transaction, prevHash string, difficulty int) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
	}

	block.Mine(difficulty)
	return block
}