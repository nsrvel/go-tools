package utils

import (
	mathRand "math/rand"
	"time"
)

var seededRand *mathRand.Rand = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

func GenerateStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
