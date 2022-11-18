package main 

import (
	"fmt"
	"os"
	"bufio"
)

func prompt(vals ...interface{}) (string, error) {
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
