package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	var sb strings.Builder
	sb.Grow(length)
	for range length {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}

func main() {
	fmt.Println(RandomString(10)) // Ví dụ: "aB3dEfGhIj"
}
