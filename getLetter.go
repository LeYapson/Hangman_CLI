package main

import (
	"fmt"
)

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