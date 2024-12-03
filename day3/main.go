package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main () {
    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error readding %s\n", os.Args[1])
    }

    lines := strings.Split(string(data), "\n") // Array of lines
    
    var ans1 int = 0
    var ans2 int = 0

    for _, line := range lines {
        fields := strings.Fields(line) // split each line by white space
    }
    fmt.Printf("Answer part 1: %v\n", ans1)
    fmt.Printf("Answer part 2: %v\n", ans2)
}


