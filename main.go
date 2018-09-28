package main

import (
	"./crypto"
	"./diffcheck"
)

func main() {
	hash1 := crypto.SeedSha1Hashs(40)
	var hash2 []string
	hash2 = append(hash1, crypto.SeedSha1Hashs(3)...)
	dc := diffcheck.New()
	dc.DiffMain(hash1, hash2)
}
