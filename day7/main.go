package main

import (
    "fmt"
    "strings"
    "os"
    "strconv"
)


func main () {

    data, err := os.ReadFile(os.Args[1])
    if err != nil{
        fmt.Println("Error reading file", os.Args[1])
        os.Exit(1)
    }

    lines := strings.Split(string(data), "\n") //Array of lines
    lines = lines[:len(lines)-1] //Cut the last empty line
    
    ans1 := 0

    for _, line := range lines {

        str_arr := strings.Split(line, " ") //Split the numbers into array

        // str_arr[0] is the result
        //slice the number out
        result_str := str_arr[0]
        result_hldr, _ := strconv.Atoi(result_str[:len(result_str)-1])
        
        fmt.Println("Expected:",result_hldr)

        //NOTE: How to add operators in between?
        // Start from str_arr[1]
        number_list := str_arr[1:]


        op_arr_list := arrange("*", "+", len(number_list) - 1)

        for i := range op_arr_list {
            var builder strings.Builder 
            op_idx := 0

            for _, num := range number_list{
                builder.WriteString(num)
                builder.WriteString(" ")
                if op_idx < len(op_arr_list[i]) {
                    builder.WriteString(string(op_arr_list[i][op_idx]))
                    builder.WriteString(" ")
                    op_idx++
                }
            }

            new_str := strings.TrimSpace(builder.String())
            new_str_arr := strings.Split(new_str, " ")
            fmt.Println(new_str_arr)

            calculation := execute(new_str_arr)
            if calculation == result_hldr {
                ans1 += calculation
                break
            }
            fmt.Println("Calculation:", calculation)
        }
    }
    
    fmt.Println("Answer 1:", ans1)

}

//helper function to arrange operators all possible ways
//return [] of operators string : "+ * + + *"
func arrange (operator_1 string, operator_2 string, num_of_pos int) []string {
    
    num_of_arrage := 1 << num_of_pos // 2^num_of_pos
    
    arrangements := make([]string, 0, num_of_arrage)
    for i := 0; i < num_of_arrage; i++ {
         var builder strings.Builder
        //mapping to get posision (1 << bit) = 001, 010, 100. (exp: num of pos = 3)
         for bit := 0; bit < num_of_pos; bit++ {
             if i & (1 << bit) != 0 {
                builder.WriteString(operator_1)
             } else {
                 builder.WriteString(operator_2)
             }
         }

        arrangements = append(arrangements, builder.String())
    } 
    
    return arrangements
}


//helper function to execute an array of number and operator
func execute (array []string) int {

    var stack = make([]int,1)
    num, _ := strconv.Atoi(array[0])
    stack[0] = num

    for i := range array {

        if i % 2 != 0 {
            switch array[i] {
            case "+":
                n, _ := strconv.Atoi(array[i+1])
                temp := stack[0] + n
                stack[0] = temp
            case "*":
                n, _ := strconv.Atoi(array[i+1])
                temp := stack[0] * n
                stack[0] = temp
            }
        }
    }

    return stack[0]
}
