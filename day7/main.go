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
    
     

    for _, line := range lines {

        str_arr := strings.Split(line, " ") //Split the numbers into array

        // str_arr[0] is the result
        //slice the number out
        numeric := str_arr[0]
        result_hldr, _ := strconv.Atoi(numeric[:len(numeric)-1])
        
        fmt.Println("Expected:",result_hldr)

        //NOTE: How to add operators in between?
        // Start from str_arr[1]
        var builder strings.Builder 

        for _, num := range str_arr {
            builder.WriteString(num)

            //test operators

        }

        new_str := strings.TrimSpace(builder.String())
        new_str_arr := strings.Split(new_str, " ")
        fmt.Println("Calculation:",execute(new_str_arr))
    }
    

}

//helper function to arrange operators all possible ways
//return [] of operators string : "+ * + + *"
func arrange (num_of_op int) []string {
    for i := 0; i < num_of_op; i++ {
        
    } 

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
