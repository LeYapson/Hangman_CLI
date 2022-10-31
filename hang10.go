package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func hang10() {

    f, err := os.Open("hang10.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}