package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	// Generate pseudo random string with length

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)

	return fmt.Sprintf("%x", b)[2 : length+2]
}
