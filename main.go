package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	if x := 1; x == 1 {
		fmt.Println("x is equal to 1")
	}

	a := 10
	b := 10
	log.Println(a + b)

	var key [32]byte
	_, err := rand.Read(key[:])
	if err != nil {
		fmt.Println("error generating random key")
		return
	}
}
