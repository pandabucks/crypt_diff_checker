package main

import (
	"./crypto"
	"./diffcheck"
)

func main() {
	hashs := crypto.SeedSha1Hashs(40)
	dc := diffcheck.New()
}
