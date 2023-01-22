package algo

import (
	"math/rand"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generate a not deterministic, pseudorandom int from 0 to upperBound
func RandomInt(upperBound int) int {
	return seed.Intn(upperBound)
}