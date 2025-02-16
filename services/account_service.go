package services

import (
	"fmt"
	"math/rand"
)

// GenerateAccountNumber generates a unique account number
func GenerateAccountNumber() string {
	return fmt.Sprintf("112233%d", rand.Intn(89999)+10000) // Format: 112233XXXXX
}
