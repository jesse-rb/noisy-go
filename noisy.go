package noisy

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// Helper function
func containsRune(alphabet string, r rune) bool {
	for _, a := range alphabet {
		if a == r {
			return true
		}
	}
	return false
}

func RandomString(length int, alphabet *string) (string, error) {
	if alphabet == nil {
		v := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		alphabet = &v
	}

	alphabetLen := len(*alphabet)

	maxInt64 := math.MaxInt64
	minInt64 := math.MinInt64
	// Ensure alphabet length is not too big
	if alphabetLen < minInt64 || alphabetLen >= maxInt64 {
		return "", fmt.Errorf("alphabet length (l) must be: %v <= l < %v", minInt64, maxInt64)
	}

	maxRandomNum := big.NewInt(int64(alphabetLen) - 1)

	randomString := ""

	for i := 0; i < length; i++ {
		randomNum, err := rand.Int(rand.Reader, maxRandomNum)
		if err != nil {
			return "", err
		}

		randomString += string([]rune(*alphabet)[randomNum.Int64()])
	}

	return randomString, nil
}

func RandomImageData(width int, height int) string {
	return ""
}
