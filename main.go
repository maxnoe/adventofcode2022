package main

import (
    "log"
    "fmt"
    "github.com/maxnoe/adventofcode2022/aoc22"
)

func main() {
    input, err := aoc22.GetInput(2021, 1)
    if (err != nil) {
        log.Fatalf("ERROR: %s", err)
    }
    fmt.Println(input)
}
