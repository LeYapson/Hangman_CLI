package main

import (
	"fmt"
	"bufio"
	"os"
	//"github.com/01-edu/z01"
)

func play() {

	fmt.Println("attention la partie va commencer donc te rate pas, tu as dix essais")

	fmt.Println(PickandRandom())

	print("quelle est donc ton premier choix? : ")
	reader := bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	print("\n","tu as choisi la lettre ",input)
	

	if input == "a"{
		print("bien joué ton code fonctionne pour le moment")
	} else {
		print("sa marche pas encore ton truc la")
	}
}




