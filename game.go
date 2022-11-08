package main

import (
	"fmt"
	"bufio"
	"os"
)

func play() {

	fmt.Println("attention la partie va commencer donc te rate pas, tu as dix essais")

	fmt.Println(PickandRandom(), "\n")

	print("quelle est donc ton premier choix? : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	if input == "a" {
		print("bien joué ton code fonctionne pour le moment")
	}
}