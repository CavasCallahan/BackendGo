package services

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUuidv4() string {
	u := make([]byte, 16)
	_, err := rand.Read(u)

	if err != nil {
		return err.Error()
	}

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}
