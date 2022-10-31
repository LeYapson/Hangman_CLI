package main

import (
    "bufio"
    "log"
    "os"
)

func importtxt() []string {

    word := []string{}

    f, err := os.Open("words2.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

        word =append(word,scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}