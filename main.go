package main

import (
    "os"
    "fmt"
    "github.com/maxnoe/adventofcode2022/aoc22"
)

func main() {
    input, err := aoc22.GetInput(2021, 1)
    if (err != nil) {
        fmt.Println("ERROR:", err)
        os.Exit(1)
    }
    fmt.Println(input)
}
