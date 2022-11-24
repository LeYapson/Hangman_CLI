package main

import (
	"bufio"
	"os"
)

func ImportHang() []string { //import le hangman.txt

	var hang []string

	readFile, _ := os.Open("hangman.txt")

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		hang = append(hang, fileScanner.Text())
	}

	readFile.Close()
	return hang
}
