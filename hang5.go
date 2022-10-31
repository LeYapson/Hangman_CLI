package main

import (
    "bufio"
    "log"
    "os"
)

func hang5() []string {

    word := []string{}

    f, err := os.Open("hang5.txt")

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
    return word
}
