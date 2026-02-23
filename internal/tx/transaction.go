package tx

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type TxInput struct {
	PrevTxHash  [32]byte
	OutputIndex uint32
	Signature   []byte
	PubKey      []byte
}

type TxOutput struct {
	Value      uint64 // value in werts (smallest unit)
	PubKeyHash []byte
}

type Transaction struct {
	Version  uint32
	Inputs   []TxInput
	Outputs  []TxOutput
	LockTime uint32
}

func NewTransaction(inputs []TxInput, outputs []TxOutput) *Transaction {
	return &Transaction{
		Version:  1,
		Inputs:   inputs,
		Outputs:  outputs,
		LockTime: 0,
	}
}

func NewCoinbaseTransaction(to []byte, value uint64) *Transaction {
	input := TxInput{
		PrevTxHash:  [32]byte{},
		OutputIndex: 0xffffffff,
		Signature:   nil,
		PubKey:      nil,
	}

	output := TxOutput{
		Value:      value,
		PubKeyHash: to,
	}

	return &Transaction{
		Version:  1,
		Inputs:   []TxInput{input},
		Outputs:  []TxOutput{output},
		LockTime: 0,
	}
}

func (tx *Transaction) Serialize() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, tx.Version)

	for _, in := range tx.Inputs {
		buf.Write(in.PrevTxHash[:])
		binary.Write(buf, binary.LittleEndian, in.OutputIndex)

		binary.Write(buf, binary.LittleEndian, uint32(len(in.Signature)))
		buf.Write(in.Signature)

		binary.Write(buf, binary.LittleEndian, uint32(len(in.PubKey)))
		buf.Write(in.PubKey)
	}

	for _, out := range tx.Outputs {
		binary.Write(buf, binary.LittleEndian, out.Value)

		binary.Write(buf, binary.LittleEndian, uint32(len(out.PubKeyHash)))
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