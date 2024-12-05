package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "strconv"
)

func main () {

    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error readding %s\n", os.Args[1])
    }

    var ans1 = 0

    lines := strings.Split(string(data), "\n") // Array of lines
    lines = lines[:len(lines)-1] // Remove the last empty element created by split
    
    //Declare table to store order
    //Key: the number that must come before all the number in the []int value
    table := make(map[int][]int)

    //Function to add value to a key ( not knowing if the key already exist or not)
    add_value := func(key int, value int) {
        if _, exist := table[key]; !exist {
            //Initialize a new slice if the key doesn't exist
            table[key] = []int{}
        }
        //Append the value into the slice
        table[key] = append(table[key], value)
    }
    
    pattern := `[0-9]+\|[0-9]+`
    reg := regexp.MustCompile(pattern)
    
    for _, line := range lines {
        match_list := reg.FindAllString(line, -1)

        for _, order_str := range match_list {
            nums := strings.Split(order_str, "|")

            first_num, _ := strconv.Atoi(nums[0])
            second_num, _ := strconv.Atoi(nums[1])

            //insert to table
            add_value(first_num, second_num)
        }

    }

    fmt.Println(table)


    pattern2 := `[0-9]+,.*`
    reg2 := regexp.MustCompile(pattern2)

    for _, line := range lines {
        match_list2 := reg2.FindString(line)
        //fmt.Println(match_list2)
        
        var str_list = []string{}
        if match_list2 != "" {
            str_list = strings.Split(match_list2, ",")
        }
        //fmt.Println(str_list)
        
        flag, middle := validate(table, str_list)
        if flag && middle > -1 {
            ans1 += middle
            fmt.Println("asn: ", ans1)
        }
    }
    fmt.Println(ans1)

}

//helper function to validate a []string of numbers -> return the middle value of the validated string
func validate (table map[int][]int, str_list []string) (bool, int) {
    flag := true
    middle := -1

    int_list := make([]int, len(str_list))

    //NOTE: int_list is a int REVERSED list of str_list
    for i := 0; i < len(str_list); i++ {
        int_list[len(str_list) - 1 - i], _ = strconv.Atoi(str_list[i])
    }
    //fmt.Println(int_list)
    
    //iterate through revesed int list and check if the next element is a value of the curr in table
    //If yes -> flag it false
    for i, num := range int_list {
        if values, exist := table[num]; exist {
            for _, v := range values {
                if i+1 < len(int_list) && v == int_list[i+1] {
                    flag = false
                }
            } 

            if flag {
                middle = int_list[len(int_list)/2] 
                fmt.Println(middle)
            }
            //NOTE: PART 2 here when flag == false reorder them
            if !flag {
                middle = reorder(int_list, table)
            }

        }
    }

    return flag, middle
}

//helper function for PART2 reorder then get the middle
func reorder(int_list []int, table map[int][]int) (int) {
    var middle int

    
    return middle
}
