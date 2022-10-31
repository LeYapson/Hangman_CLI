package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func hang3() {

    f, err := os.Open("hang3.txt")

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
