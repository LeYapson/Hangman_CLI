package main

import (
	"fmt"
)

func getLetter(found []string) (string, error) {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for true {
		letter, err := prompt("ecrit zebi", join(found, " "))
		if err != nil {
			return "", err
		}
		if len(letter) == 1 && containsAny(alphabet, []string{letter}) {
			return letter, nil
		}
		fmt.Println("ici c la france si tu l aime pas retourne chez toi ")
	}
	return "", nil
}
