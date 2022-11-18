package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	words := ImportTxt()
	t := time.Now()
	rand.Seed(t.UnixNano())
	idx := rand.Intn(len(words))
	word := words[idx]
	nguesses := 10
	found := []string{}
	for i := 0; i < len(word); i++ {
		found = append(found, "_")
	}
	for nguesses > 0 {
		fmt.Println("You have", nguesses, "remaining guesses.")
		letter, err := getLetter(found)
		if err != nil {
			fmt.Println("error reading from console")
			return
		}
		if !containsAny(word, []string{letter}) {
			nguesses--
		}
		if updateFound(found, word, letter) {
			fmt.Println("you win! the word was:", word)
			return
		}
	}
	fmt.Println("you lose! the word was:", word)
}
