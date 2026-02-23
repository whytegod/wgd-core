package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"wgd-core/internal/tx"
)

type BlockHeader struct {
	Version    uint32
	PrevHash   [32]byte
	MerkleRoot [32]byte
	Timestamp  int64
	Bits       uint32
	Nonce      uint32
}

type Block struct {
	Header       BlockHeader
	Transactions []*tx.Transaction
	Height       uint64
}

func NewBlock(prevHash [32]byte, transactions []*tx.Transaction, height uint64, bits uint32) *Block {
	block := &Block{
		Header: BlockHeader{
			Version:   1,
			PrevHash:  prevHash,
			Timestamp: time.Now().Unix(),
			Bits:      bits,
			Nonce:     0,
		},
		Transactions: transactions,
		Height:       height,
	}

	block.Header.MerkleRoot = CalculateMerkleRoot(transactions)
	return block
}

func NewGenesisBlock() *Block {
	var emptyHash [32]byte
	return NewBlock(emptyHash, []*tx.Transaction{}, 0, 0x1d00ffff)
}

func (b *Block) Hash() [32]byte {
	return b.Header.Hash()
}

func (h *BlockHeader) Hash() [32]byte {
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, h.Version)
	buffer.Write(h.PrevHash[:])
	buffer.Write(h.MerkleRoot[:])
	binary.Write(&buffer, binary.LittleEndian, h.Timestamp)
	binary.Write(&buffer, binary.LittleEndian, h.Bits)
	binary.Write(&buffer, binary.LittleEndian, h.Nonce)

	first := sha256.Sum256(buffer.Bytes())
	second := sha256.Sum256(first[:])

	return second
}

func CalculateMerkleRoot(transactions []*tx.Transaction) [32]byte {
	var hashes [][]byte

	if len(transactions) == 0 {
		return sha256.Sum256([]byte{})
	}

	for _, tx := range transactions {
		txHash := tx.Hash()
		hashCopy := make([]byte, 32)
		copy(hashCopy, txHash[:])
		hashes = append(hashes, hashCopy)
	}

	for len(hashes) > 1 {
		var newLevel [][]byte

		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				hashes = append(hashes, hashes[i])
			}

			combined := append(hashes[i], hashes[i+1]...)
			hash := sha256.Sum256(combined)
			newLevel = append(newLevel, hash[:])
		}

		hashes = newLevel
	}

	var merkleRoot [32]byte
	copy(merkleRoot[:], hashes[0])
	return merkleRoot
}