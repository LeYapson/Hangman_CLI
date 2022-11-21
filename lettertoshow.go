package main

import (
	"math/rand"
	"time"
)

func LettertoShow(s string) string { //aficher les lettre dans le int
	rand.Seed(time.Now().UnixNano())
	inRune := []rune(s)
	var pick rune
	randomIndex := rand.Intn(len(inRune))
	pick = inRune[randomIndex]
	print(string(pick))

	return string(pick)
}
