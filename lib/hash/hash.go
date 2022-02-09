package hash

import (
	"math/rand"
	"strings"
	"time"
)

const HASH_BASE = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func CreateSixDigitHash() string {
	var hashes = make([]string, 6)
	for i := 0; i < 6; i++ {
		hashes[i] = chooseRandomChar()
	}
	hashStr := strings.Join(hashes, "")
	return hashStr
}

func chooseRandomChar() string {
	rand.Seed(time.Now().UnixNano())
	randomDigit := string(HASH_BASE[rand.Intn(62)])
	return randomDigit
}
