package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(vals ...interface{}) (string, error) { //permet de rentrer une valeur au terminal
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
