package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
)

func main () {
    var list_1 []int
    var list_2 []int
    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error readding %s\n", os.Args[1])
    }

    lines := strings.Split(string(data), "\n") // Array of lines
    
    for _, line := range lines {
        fields := strings.Fields(line) // split each line by white space

        //Parse each field into integer and append to the corresponding list
        for i, field := range fields {
            num, err := strconv.Atoi(field)
            if err != nil {
                fmt.Printf("Cannot convert string to number\n")
                return
            }
            switch i {
                case 0: 
                    list_1 = append(list_1,num)
                case 1:
                    list_2 = append(list_2,num)
            }
        }
    }

    //pre-built map for Part 2
    freq_list_2 := make(map[int]int)
    for i := 0; i < len(list_2); i++ {
        freq_list_2[list_2[i]]++
    }

    // Main logic
    // Part 1
    sort.Ints(list_1)
    sort.Ints(list_2)

    var ans1 = 0
    var ans2 = 0

    for i := 0; i < len(list_1); i++ {
        //Part 1
        sub := list_1[i] - list_2[i]
        if sub < 0 {
            ans1 -= sub
        } else {
            ans1 += sub
        }

        //Part 2
        ans2 += list_1[i]*freq_list_2[list_1[i]]
    }

    fmt.Printf("Result Part 1: %v\n", ans1)
    fmt.Printf("Result Part 2: %v\n", ans2)
}
