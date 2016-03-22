package tools

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
)

// ErrInvalidMax is returned when an int64 less than 1 is specified
// for a random integer range of 0 to max.
var ErrInvalidMax = errors.New("max cannot be less than 1")

// CalcEntropy returns the calculated entropy where "t" is the total number of
// items available for random selection, and "n" is the number of random
// items selected.
func CalcEntropy(t, n float64) float64 {
	return math.Log2(math.Pow(t, n))
}

// RandomInt64 generates a random 64-bit integer from 0 to "max."
func RandomInt64(max int64) (int64, error) {
	if max < 1 {
		return 0, ErrInvalidMax
	}

	i, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}

	return i.Int64(), nil
}
