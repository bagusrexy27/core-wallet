package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func CalculateChecksum(userId string, balance int64) string {
	secret := os.Getenv("CHECKSUM_SECRET")
	if secret == "" {
		secret = "secret"
	}

	data := fmt.Sprintf("%s:%d:%s", userId, balance, secret)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func ValidateChecksum(userId string, balance int64, checksum string) bool {
	expectedChecksum := CalculateChecksum(userId, balance)
	return expectedChecksum == checksum
}
