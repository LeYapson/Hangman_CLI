package main

import (
	"fmt"
	"math/rand"
	"time"
	//"github.com/01-edu/z01"
)

func PickandRandom() {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano())
	wordtofind := ImportTxt()
	word := (wordtofind[rand.Intn(len(wordtofind))])
	Blind(word)
	//n := len(wordtofind) / 2 - 1
	
fmt.Println(Blind(word))
}
