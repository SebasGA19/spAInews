package common

import (
	"crypto/rand"
	"encoding/hex"
)

func RandString(length int) string {
	buffer := make([]byte, length)
	_, readError := rand.Read(buffer)
	if readError != nil {
		panic(readError)
	}
	return hex.EncodeToString(buffer)
}
