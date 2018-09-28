package crypto

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// create random SHA1 map object .
func SeedCryptos(num int) []string {
	h := sha1.New()
	var result []string
	for i := 0; i < num; i++ {
		random_letter := RandString(100)
		h.Write([]byte(random_letter))
		bs := h.Sum(nil)
		result = append(result, hex.EncodeToString(bs))
	}
	return result
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}
