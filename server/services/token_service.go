package services

import (
	"math/rand"
)

func GenerateStaticToken() string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 128)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)

}
