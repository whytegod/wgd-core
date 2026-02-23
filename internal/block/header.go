package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type BlockHeader struct {
	Version    uint32
	PrevHash   [32]byte
	MerkleRoot [32]byte
	Timestamp  uint32
	Bits       uint32
	Nonce      uint32
}

func (h *BlockHeader) Serialize() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, h.Version)
	buf.Write(h.PrevHash[:])
	buf.Write(h.MerkleRoot[:])
	binary.Write(buf, binary.LittleEndian, h.Timestamp)
	binary.Write(buf, binary.LittleEndian, h.Bits)
	binary.Write(buf, binary.LittleEndian, h.Nonce)

	return buf.Bytes()
}

func (h *BlockHeader) Hash() [32]byte {
	first := sha256.Sum256(h.Serialize())
	second := sha256.Sum256(first[:])
	return second
}