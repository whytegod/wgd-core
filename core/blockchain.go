package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents one block in the chain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain is a collection of blocks
type Blockchain struct {
	Blocks []Block
}

// Create a new block
func createBlock(index int, data string, prevHash string) Block {
	timestamp := time.Now().String()

	record := strconv.Itoa(index) + timestamp + data + prevHash
	hash := sha256.Sum256([]byte(record))

	return Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevHash,
		Hash:      hex.EncodeToString(hash[:]),
	}
}

// Create genesis block
func createGenesisBlock() Block {
	return createBlock(0, "Genesis Block", "")
}

// Initialize new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []Block{createGenesisBlock()},
	}
}

// Add new block
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := createBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}