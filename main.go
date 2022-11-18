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
		fmt.Println("Tu as", nguesses, "essais pour réussir")
		letter, err := getLetter(found)
		if err != nil {
			fmt.Println("ERREUR SUR LA CONSOLE")
			return
		}
		if !containsAny(word, []string{letter}) {
			nguesses--
			if nguesses == 9 {
				printhang0()
			} else if nguesses == 8 {
				printhang1()
			} else if nguesses == 7 {
				printhang2()
			} else if nguesses == 6 {
				printhang3()
			} else if nguesses == 5 {
				printhang4()
			} else if nguesses == 4 {
				printhang5()
			} else if nguesses == 3 {
				printhang6()
			} else if nguesses == 2 {
				printhang7()
			} else if nguesses == 1 {
				printhang8()
			} else if nguesses == 0 {
				printhang9()
			}

		}
		if updateFound(found, word, letter) {
			fmt.Println("gg mtn bouge, le mot c'était bien ", word)
			return
		}
	}
	fmt.Println("ff la prochaine fois, le mot c'était", word, "try again (sale merde)")
}
