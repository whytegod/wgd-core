package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents a single block in the chain
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// calculateHash creates SHA256 hash of block contents
func calculateHash(b Block) string {
	record := string(rune(b.Index)) + b.Timestamp + b.Data + b.PreviousHash
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// NewBlock creates a new block
func NewBlock(index int, data string, prevHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().Format(time.RFC3339),
		Data:         data,
		PreviousHash: prevHash,
	}

	block.Hash = calculateHash(block)
	return block
}