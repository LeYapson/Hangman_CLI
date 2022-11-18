package main

import (
	"fmt"
)

func getLetter(found []string) (string, error) {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for true {
		letter, err := input("écrit une lettre", join(found, " "))
		if err != nil {
			return "", err
		}
		if len(letter) == 1 && containsAny(alphabet, []string{letter}) {
			return letter, nil
		}
		fmt.Println("ERREUR! tu ne peut écrire que des lettres minuscules")
	}
	return "", nil
}
