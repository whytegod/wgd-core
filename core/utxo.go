package core

type UTXOSet struct {
	UTXOs map[string][]TXOutput
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		UTXOs: make(map[string][]TXOutput),
	}
}