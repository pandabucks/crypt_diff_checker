package main

import (
	"./crypto"
	"fmt"
)

func main() {
	data := crypto.SeedCryptos(40)
	fmt.Println(data)
}
