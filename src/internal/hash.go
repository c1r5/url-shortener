package internal

import (
	"crypto/sha256"
	"encoding/binary"
)

func Hash(s string) uint64 {
	hash := sha256.Sum256([]byte(s))

	// take first 8 bytes → uint64
	return binary.BigEndian.Uint64(hash[:8])
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeBase62(n uint64) string {
	if n == 0 {
		return string(charset[0])
	}

	var result []byte
	base := uint64(len(charset))

	for n > 0 {
		result = append([]byte{charset[n%base]}, result...)
		n /= base
	}

	return string(result) // return only the first 6 characters
}
