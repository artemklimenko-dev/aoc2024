package main

import (
    "fmt"
    "os"
    "github.com/artemklimenko-dev/aoc2024/days/day1"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <day>")
        return
    }
    day := os.Args[1]
	
    switch day {
    case "1":
        day1.Solve()
    default:
        fmt.Println("Day not present yet.")
    }
}