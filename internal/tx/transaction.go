package tx

import (
	"crypto/sha256"
	"encoding/binary"
	"bytes"
)

type TxInput struct {
	PrevTxHash [32]byte
	OutputIndex uint32
	Signature []byte
	PubKey    []byte
}

type TxOutput struct {
	Value uint64 // in werts
	PubKeyHash []byte
}

type Transaction struct {
	Version  uint32
	Inputs   []TxInput
	Outputs  []TxOutput
	LockTime uint32
}

func (tx *Transaction) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, tx.Version)

	for _, in := range tx.Inputs {
		buf.Write(in.PrevTxHash[:])
		binary.Write(buf, binary.LittleEndian, in.OutputIndex)
		buf.Write(in.Signature)
		buf.Write(in.PubKey)
	}

	for _, out := range tx.Outputs {
		binary.Write(buf, binary.LittleEndian, out.Value)
		buf.Write(out.PubKeyHash)
	}

	binary.Write(buf, binary.LittleEndian, tx.LockTime)

	return buf.Bytes()
}

func (tx *Transaction) Hash() [32]byte {
	first := sha256.Sum256(tx.Serialize())
	second := sha256.Sum256(first[:])
	return second
}