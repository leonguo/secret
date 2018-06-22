package util

import "math/rand"

func RandomSample(letters string, n int) string {
	b := make([]byte, n)
	lenLetter := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(lenLetter)]
	}
	return string(b)
}