package main

import (
	"math/rand"
	"time"
)

func randomWord([]string) string {
	words := ImportTxt()
	t := time.Now()
	rand.Seed(t.UnixNano())
	idx := rand.Intn(len(words))
	word := words[idx]
	return word
}
