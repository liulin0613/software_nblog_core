package utils

import (
	"math/rand"
	"time"
)

const (
	letters = "QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm"
	number  = "0123456789"
)

func GenerateRandomNumber(start int, end int) int {
	if end < start {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(end-start) + start
}

func Random(n int, chars string) string {
	if n < 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n, n)
	l := len(chars)
	for i := 0; i < n; i++ {
		b[i] = chars[r.Intn(l)]
	}
	return string(b)
}

func RandomString(n int) string {
	return Random(n, letters)
}

func RandomNumber(n int) string {
	return Random(n, number)
}
