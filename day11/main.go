package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
    
    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Cannot open file")
        os.Exit(-1)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    elements := strings.Split(lines[0], " ")

    var nums []int
    for _, element := range elements {
        num, _ := strconv.Atoi(element)
        nums = append(nums, num)
    }

    //Recursion
    //each stone will go its' own path
    //count total after x split

	// Memoization map
	memo1 := make(map[string]int)
	memo2 := make(map[string]int)
    
    ans1 := 0
    ans2 := 0
    for _, num := range nums {
        ans1 += count(num, 25, memo1)
        ans2 += count(num, 75, memo2)
    }
    
    fmt.Println("Answer 1:",ans1)
    fmt.Println("Answer 2:",ans2)

}

func count(num int, split int, memo map[string]int) int {
	// Create a unique key for memoization
	key := fmt.Sprintf("%d-%d", num, split)
	if val, found := memo[key]; found {
		return val
	}

	if split == 0 {
		return 1
	}

	var result int
	if num == 0 {
		result = count(1, split-1, memo)
	} else if checkEven(num) {
		num1, num2 := splitNum(num)
		result = count(num1, split-1, memo) + count(num2, split-1, memo)
	} else {
		result = count(num*2024, split-1, memo)
	}

	// Store result in memo
	memo[key] = result
	return result
}

func checkEven(num int) bool {
    str := strconv.Itoa(num)
    if len(str) % 2 == 0 {
        return true
    } else {
        return false
    }
}

func splitNum(num int) (int, int) {
    str := strconv.Itoa(num)
    if len(str) % 2 != 0 {
        panic("Cannot split odd number")
    }
    
    num1, _ := strconv.Atoi(str[:len(str)/2])
    num2, _ := strconv.Atoi(str[len(str)/2:])

    return num1, num2
}

