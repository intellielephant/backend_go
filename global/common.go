package global

import (
	"math/rand"
)

func RandomString(n int) string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for i := range result {
		randomByteIndex := rand.Intn(len(letters))
		result[i] = letters[randomByteIndex]
	}
	return string(result)
}

func GenInviteCode() string {
	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, 6)
	for i := range result {
		randomByteIndex := rand.Intn(len(letters))
		result[i] = letters[randomByteIndex]
	}
	return string(result)
}

