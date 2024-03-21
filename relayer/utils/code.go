package utils

import (
	"math/rand"
	"time"
)

var b58Alphabet = []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func GenerateCode(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = b58Alphabet[rand.Int()%len(b58Alphabet)]
	}
	idResult := string(id)
	return idResult
}
