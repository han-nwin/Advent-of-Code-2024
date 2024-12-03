package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main () {

    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error readding %s\n", os.Args[1])
    }

    memory := (string(data)) // turn data into a string

    //regex to get a []string of mul(1,2)
    pattern := `mul\([0-9][0-9]*,[0-9][0-9]*\)`
    reg, _:= regexp.Compile(pattern)
    match_list := reg.FindAllString(memory, -1) //-1: find all string -> return []string

    var ans1 float64 = 0
    //iterate through the []string
    for _, element := range match_list {
        //Regex to get each number
        num_pattern := `[0-9][0-9]*`
        num_reg, _ := regexp.Compile(num_pattern)
        number_list := num_reg.FindAllString(element, -1)

        var product float64 = 1
        for _, num := range number_list {
            num_int, _ := strconv.Atoi(num)
            product *= float64(num_int)
        }
        ans1 += product
    }

    //Part 2
    pattern2 := `mul\([0-9][0-9]*,[0-9][0-9]*\)|do\(\)|don't\(\)`
    reg2, _:= regexp.Compile(pattern2)
    match_list2 := reg2.FindAllString(memory, -1) //-1: find all string -> return []string

    var ans2 float64 = 0
    flag := true
    //iterate through the []string
    for _, element := range match_list2 {
        if element == "do()" {
            flag = true
            continue
        }
        if element == "don't()" {
            flag = false
            continue
        }
        if flag == false {
            continue
        }

        //Regex to get each number
        num_pattern2 := `[0-9][0-9]*`
        num_reg2, _ := regexp.Compile(num_pattern2)
        number_list2 := num_reg2.FindAllString(element, -1)

        var product2 float64 = 1
        for _, num2 := range number_list2 {
            num_int2, _ := strconv.Atoi(num2)
            product2 *= float64(num_int2)
        }
        ans2 += product2
    }
    fmt.Printf("Answer part 1: %v\n", ans1)
    fmt.Printf("Answer part 2: %v\n", ans2)
}


