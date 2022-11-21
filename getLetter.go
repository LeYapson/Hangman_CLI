package main

import (
	"fmt"
)

func getLetter(found []string) (string, error) {     //verifie si la valeur de input et bien une lettre

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
