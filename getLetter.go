package main

import (
	"fmt"
)

func getLetter(found []string) (string, error) {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for true {
		letter, err := prompt("stp choisie une lettre", join(found, " "))
		if err != nil {
			return "", err
		}
		if len(letter) == 1 && containsAny(alphabet, []string{letter}) {
			return letter, nil
		}
		fmt.Println("entrée invalide: il n'y a que des lettres minuscules.")
	}
	return "", nil
}