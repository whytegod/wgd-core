package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

const Difficulty = 4

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []*Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

func NewBlock(transactions []*Transaction, prevHash string, index int) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}
	block.mine()
	return block
}

func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) +
		strconv.FormatInt(b.Timestamp, 10) +
		b.PrevHash +
		strconv.Itoa(b.Nonce)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (b *Block) mine() {
	prefix := strings.Repeat("0", Difficulty)

	for {
		hash := b.calculateHash()
		if strings.HasPrefix(hash, prefix) {
			b.Hash = hash
			break
		}
		b.Nonce++
	}
}