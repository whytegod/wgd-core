package core

import (
	"bytes"
	"crypto/sha3"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []*Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int
	Difficulty   int
}

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

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	_ = encoder.Encode(b)
	return res.Bytes()
}

func (b *Block) HashBlock() []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(int64(b.Index)),
			IntToHex(b.Timestamp),
			b.PrevHash,
			b.HashTransactions(),
			IntToHex(int64(b.Nonce)),
		},
		[]byte{},
	)

	hash := sha3.Sum256(data)
	return hash[:]
}

func (b *Block) Mine() {
	target := bytes.Repeat([]byte{0}, b.Difficulty)

	for {
		hash := b.HashBlock()

		if bytes.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}

	fmt.Printf("Block %d mined: %x\n", b.Index, b.Hash)
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	tree := NewMerkleTree(txHashes)
	return tree.RootNode.Data
}