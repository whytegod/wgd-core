package wallet

import (
	"bytes"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
)

const VersionByte = byte(0x49) // WGD network prefix (example)

func Checksum(payload []byte) []byte {
	first := sha256.Sum256(payload)
	second := sha256.Sum256(first[:])
	return second[:4]
}

func PubKeyHashToAddress(pubKeyHash []byte) string {
	versionedPayload := append([]byte{VersionByte}, pubKeyHash...)
	checksum := Checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)
	return base58.Encode(fullPayload)
}