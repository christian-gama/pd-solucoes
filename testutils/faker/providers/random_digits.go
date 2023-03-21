package providers

import (
	"crypto/rand"
	"math/big"
)

func generateRandomDigits(length int) []int {
	digits := make([]int, length)
	const maxInt = 10

	for i := range digits {
		n, err := rand.Int(rand.Reader, big.NewInt(maxInt))
		if err != nil {
			panic(err)
		}
		digits[i] = int(n.Int64())
	}

	return digits
}
