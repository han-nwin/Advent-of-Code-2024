package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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
    ans2 := 0

    var mutex sync.Mutex //mutex to protect ans1 ans2
    var wg sync.WaitGroup //synchronize var

    for _, line := range lines {
        wg.Add(1) //increment Waitgroup counter
        go func(line string) {
            defer wg.Done()//Decrement counter when goroutine done
            

            str_arr := strings.Split(line, " ") //Split the numbers into array

            // str_arr[0] is the result
            //slice the number out
            result_str := str_arr[0]
            result_hldr, _ := strconv.Atoi(result_str[:len(result_str)-1])
            

            //NOTE: How to add operators in between?
            // Start from str_arr[1]
            number_list := str_arr[1:]


            //Create operators array and arrange them
            ops := []string{"*", "+"}
            op_arr_list := arrange(ops, len(number_list) - 1)

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

                calculation := execute(new_str_arr)
                if calculation == result_hldr {
                    mutex.Lock() //Synchorize total ans here
                    ans1 += calculation
                    mutex.Unlock()
                    break
                }
            }

            //NOTE: PART 2, HOW TO ADD || HERE? 
            
            //new operators array
            ops2 := []string{"*", "+", "||"}
            op_arr_list2 := arrange(ops2, len(number_list) - 1)

            for i := range op_arr_list2 {
                var builder strings.Builder 
                op_idx := 0

                for _, num := range number_list{
                    builder.WriteString(num)
                    builder.WriteString(" ")
                    if op_idx < len(op_arr_list2[i]) {
                        builder.WriteString(string(op_arr_list2[i][op_idx]))
                        builder.WriteString(" ")
                        op_idx++
                    }
                }

                new_str := strings.TrimSpace(builder.String())
                new_str_arr := strings.Split(new_str, " ")

                calculation := execute(new_str_arr)
                if calculation == result_hldr {
                    mutex.Lock() //Synchorize total ans here
                    ans2 += calculation
                    mutex.Unlock()
                    break
                }
            }

        }(line)
    }

    wg.Wait() //wait for all goroutine to complete
    
    fmt.Println("Answer 1:", ans1)
    fmt.Println("Answer 2:", ans2)

}

//helper function to arrange operators all possible ways
//return [] of operators string : "+ * + + *"
func arrange (operators []string, num_of_pos int) [][]string {
    
    num_of_arrage := 1 
    for i := 0; i < num_of_pos; i++ {
        num_of_arrage *= len(operators) // operators^num_of_pos. Example: 3 operator, 5 pos = 3^5 arrangements
    }
    
    arrangements := make([][]string, 0, num_of_arrage)

    for i := 0; i < num_of_arrage; i++ {
        arrangement := make([]string, num_of_pos)
        base := i

        //Generate combination here
        for pos := 0; pos < num_of_pos; pos++ {
            arrangement[pos] = operators[base % len(operators)]
            base /= len(operators)
        }
        arrangements = append(arrangements, arrangement)
    } 
    
    return arrangements
}

//helper function to execute an array of number and operator
func execute (array []string) int {

    num, _ := strconv.Atoi(array[0])
    holder := num

    for i := 1; i < len(array); i += 2 {
        n, _ := strconv.Atoi(array[i+1])
        switch array[i] {
        case "+":
            holder += n
        case "*":
            holder *= n
        case "||":
            holder = holder*(int(math.Pow(10, float64(count_digits(n))))) + n
        }
    }
    return holder
}

//helper function to count digit
func count_digits(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}
