package main

import (
	"bufio"
	"os"
)

func Input(str string) string {
	print(str)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}
	
	