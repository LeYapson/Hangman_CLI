package main

import (
	//"fmt"
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
	for i := 0; i < len(array); i++{
		if array[i] >= 'a' && array[i] <= 'z' {
			
			array[i] = '_'
			
		}
		z01.PrintRune(array[i])
	}
	return string(array)
	
	
	
	

}

func LettertoShow(s string) string {
	rand.Seed(time.Now().UnixNano())
	letter := (s)
	var lettertoprint byte
	for i:=0; i<=len(letter)/2-1; i++ {
		lettertoprint = (letter[rand.Intn(len(letter))])
	fmt.Print(string(lettertoprint))
	}
	return string(lettertoprint)
}
