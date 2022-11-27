package main

import (
    "fmt"
    "github.com/maxnoe/adventofcode2022/aoc22"
)

func main() {
    input, err := aoc22.GetInput(2021, 1)
    if (err != nil) {
        fmt.Println("ERROR:", err)
        return
    }
    fmt.Println(input)
}
