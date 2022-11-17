package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func containsAny(s string, chars []string) bool {
	for _, ch := range s {
		for _, ch2 := range chars {
			if string(ch) == ch2 {
				return true
			}
		}
	}
	return false
}

func join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	s := ""
	lastIdx:= len(strings) - 1
	for _, v := range strings[0:lastIdx] {
		s += v + separator    // s = s + v + separator
	}
	return s + strings[lastIdx]
}

func getLetter(found []string) (string, error) {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for true {
		letter, err := prompt("pick a letter", join(found, " "))
		if err != nil {
			return "", err
		}
		if len(letter) == 1 && containsAny(alphabet, []string{letter}) {
			return letter, nil
		}
		fmt.Println("Invalid input: must enter a single lowercase letter.")
	}
	return "", nil
}

func updateFound(found []string, word string, letter string) bool {
	complete := true
	for i, r := range word {
		if letter == string(r) {
			found[i] = letter
		}
		if found[i] == "_" {
			complete = false
		}
	}
	return complete
}




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


func prompt(vals ...interface{}) (string, error) {
	if len(vals) != 0 {
		fmt.Println(vals...)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
