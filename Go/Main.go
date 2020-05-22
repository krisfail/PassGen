package main

import (
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	clipboard.WriteAll(randomText(16))
}

func randomText(length int) string {
	const charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}
