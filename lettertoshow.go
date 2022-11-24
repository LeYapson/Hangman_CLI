package main

import (
	"math/rand"
	"time"
)

func LettertoShow(s string) string { //aficher les lettre dans le int
	rand.Seed(time.Now().UnixNano())
	inRune := []rune(s)
	var pick rune
	l := len(inRune)/2 - 1
	for i := 0; i < l; i++ {

		randomIndex := rand.Intn(len(inRune))
		pick = inRune[randomIndex]
		if pick != pick+1 {
			print(string(pick))
		}

	}
	return string(pick)
}
