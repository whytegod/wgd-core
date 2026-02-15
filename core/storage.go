package core

import (
	"encoding/json"
	"os"
)

func (bc *Blockchain) SaveToFile(filename string) {
	data, _ := json.MarshalIndent(bc, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func LoadBlockchain(filename string) *Blockchain {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	var bc Blockchain
	json.Unmarshal(data, &bc)

	return &bc
}