package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Reward    float64
}

func NewBlock(index int, prevHash string, data string, reward float64) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
		Reward:    reward,
	}

	block.Hash = calculateHash(block)
	return block
}

func calculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}