package genesis

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

const (
	GenesisMessage = "23 Feb 2026 — WGD Genesis — A Neutral Global Monetary Protocol"
	GenesisTime    = int64(1771804800) // 23 Feb 2026 UTC timestamp (fixed)
)

type GenesisBlock struct {
	Version    int
	Timestamp  int64
	PrevHash   string
	MerkleRoot string
	Nonce      uint64
	Hash       string
}

func CreateGenesisBlock() *GenesisBlock {
	version := 1
	prevHash := "0"
	nonce := uint64(0)

	merkleRoot := calculateMerkleRoot([]byte(GenesisMessage))

	blockHeader := buildHeader(version, GenesisTime, prevHash, merkleRoot, nonce)

	hash := calculateHash(blockHeader)

	return &GenesisBlock{
		Version:    version,
		Timestamp:  GenesisTime,
		PrevHash:   prevHash,
		MerkleRoot: merkleRoot,
		Nonce:      nonce,
		Hash:       hash,
	}
}

func calculateMerkleRoot(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func buildHeader(version int, timestamp int64, prevHash, merkleRoot string, nonce uint64) []byte {
	record := string(rune(version)) +
		string(rune(timestamp)) +
		prevHash +
		merkleRoot +
		string(rune(nonce))

	return []byte(record)
}

func calculateHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}