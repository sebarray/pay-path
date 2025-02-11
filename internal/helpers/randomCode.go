package helpers

import (
	"crypto/rand"
	"math/big"
)

const charset = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"

func RandomAlphanumeric(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}
