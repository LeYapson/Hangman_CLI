package main

import (
	"fmt"
	"math/rand"
	"time"
)

func LettertoShow(s string) string {
	rand.Seed(time.Now().UnixNano())
	inRune := []rune(s)
	var pick rune
	l := len(inRune)/2 - 1
	for i := 0; i < l; i++ {

		randomIndex := rand.Intn(len(inRune))
		pick = inRune[randomIndex]
		fmt.Print(string(pick))
	}

	return string(pick)
}
