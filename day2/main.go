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
        valid := check_valid(fields) //type []string
         
        if valid {//Part 1
            ans1++
            ans2++
        } else { //Part 2
            for i := 0; i < len(fields); i++ {
                //make a new slice without index i
                new_fields := make([]string, len(fields)-1)
                copy(new_fields, fields[:i])
                copy(new_fields[i:], fields[i+1:])
                if check_valid(new_fields) {
                    ans2++
                    break
                }
            }
        }

    }
    fmt.Printf("Answer part 1: %v\n", ans1)
    fmt.Printf("Answer part 2: %v\n", ans2)
}

func check_valid(fields []string) (bool) {

    var flag = false // false: increasing, true: decreasing

    if len(fields) < 2 {
        return false
    }

    var first_num, _ = strconv.Atoi(fields[0])
    var second_num, _ = strconv.Atoi(fields[1])
    if (first_num > second_num) {
        flag = true
    } else if first_num == second_num {
        return false
    }
    
    var valid = true //flag to check the sequence logic
    for i := 1; i < len(fields); i++ {
        curr, _ := strconv.Atoi(fields[i])
        prev,_ := strconv.Atoi(fields[i-1])

        if flag { //decreasing
            if curr >= prev || prev - curr > 3 {
                valid = false
                break
            }
        } else { //increasing
            if curr <= prev || curr - prev > 3 {
                valid = false
                break
            }
        }
    }
        
    if valid {
        return true
    } else {
        return false
    }

}
