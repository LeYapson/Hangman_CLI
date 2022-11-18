package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/01-edu/z01"
)

func PickandRandom() string {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano())
	wordtofind := ImportTxt()
	word := (wordtofind[rand.Intn(len(wordtofind))])
	array := []rune(word)
	fmt.Print(word)

	for i := 0; i < len(array); i++ {
		if array[i] >= 'a' && array[i] <= 'z' {

			array[i] = '_'

		}
		z01.PrintRune(array[i])

	}
	LettertoShow(word)
	return string(array)

}
