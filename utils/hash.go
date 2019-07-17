package Utils

import (
	blake2s "golang.org/x/crypto/blake2s"
)

func GetHash(key []byte) ([]byte, error) {
	hash := blake2s.Sum256(key)
	return hash[:], nil
}
