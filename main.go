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
	if nguesses == 10 { //indice
		fmt.Println("INDICE: les lettre suivantes sont présentes dans le mot :", "\n")
		LettertoShow(word)
		fmt.Println("\n")
	}
	for nguesses > 0 { //nbre essais + input
		fmt.Println("Tu as", nguesses, "essais pour réussir")
		letter, err := getLetter(found)
		if err != nil {
			fmt.Println("ERREUR SUR LA CONSOLE")
			return
		}
		if !containsAny(word, []string{letter}) { // aficher le hangman
			nguesses--
			PrintHang(nguesses)

		}
		if updateFound(found, word, letter) {
			fmt.Println("Bravo, tu as gagné, le mot était bien :", word) //fin du jeu
			return
		}
	}
	fmt.Println("ff la prochaine fois, le mot c'était :", word, "tu peux toujours retenter ta chance")
}
