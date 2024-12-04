
package main

import (
    "fmt"
    "os"
    "regexp"
    //"strconv"
    "strings"
)

func main () {

    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error readding %s\n", os.Args[1])
    }

    var max_len = 0

    var ans1 = 0

    //Process line
    lines := strings.Split(string(data), "\n") // Array of lines
    lines = lines[:len(lines)-1] // Remove the last empty element created by split
    for i, line := range lines {
        fmt.Printf("Line %d: '%s' (length: %d)\n", i, line, len(line))
        ans1 += find_xmas(line)
        max_len++//get number of rows to use later
    }

    //Process reversed line
    for _, line := range lines {
        reversed_line := reverse(line)
        ans1 += find_xmas(reversed_line)
    }


    //Process column
    var columns = make([]string, max_len)

    for _, line := range lines {
        for i, char := range line {
            columns[i] += string(char)
        }
    }
    for _, column := range columns {
        ans1 += find_xmas(column)
    }

    //Process column reverse
    for _, column := range columns {
        reversed_column := reverse(column)
        ans1 += find_xmas(reversed_column)
    }

    //Process diag. There are 2 diag top left to bottom right, top right to bottom left
    max_2_len := len(lines) 
    if max_2_len < len(columns) {
        max_2_len = len(columns)
    }
    var diags1 = make([]string, max_2_len*2 - 1)
    var diags2 = make ([]string, max_2_len*2 - 1)
    var diags1_idx = 0
    var diags2_idx = 0

    //building each string
    for row := 0; row < len(lines); row++ {
        for col, char := range lines[row] {
			// Top left to bottom right
			diags1_idx = row + col
			diags1[diags1_idx] += string(char)

			//Top right to bottom left
            diags2_idx = row - col + (len(columns)- 1)
			diags2[diags2_idx] += string(char)
        }
    }
    
    //diag 1
    for _, diag := range diags1 {
        ans1 += find_xmas(diag)
    }

    //Process diag reverse
    for _, diag := range diags1 {
        reversed_diag := reverse(diag)
        ans1 += find_xmas(reversed_diag)
    }

    //diag 2
    for _, diag := range diags2 {
        ans1 += find_xmas(diag)
    }

    //Process diag reverse
    for _, diag := range diags2 {
        reversed_diag := reverse(diag)
        ans1 += find_xmas(reversed_diag)
    }

    //Final process to count number of XMAS
    fmt.Println(ans1)

    var ans2 = 0

   //Part 2 row 
    for i := 0; i < len(lines) - 2; i++ {
        pattern1 := `M.S|S.M`
        pattern2 := `.A.`

        indices1 := findMatchStarts(pattern1, lines[i]) // M.S
        indices2 := findMatchStarts(pattern2, lines[i+1]) // .A.
        indices3 := findMatchStarts(pattern1, lines[i+2])// M.S
        fmt.Println(indices1,"-----")
        fmt.Println(indices2,"-----")
        fmt.Println(indices3,"-----")
        for _, match1 := range indices1 {
            for _, match2 := range indices2 {
                for _, match3 := range indices3 {
                    if match3 == match2 && match2 == match1 {
                        ans2++
                    }
                }
            }
        }
        fmt.Println(ans2)
    }

    //Part 2 column
    for i := 0; i < len(columns) - 2; i++ {
        pattern1 := `M.S|S.M`
        pattern2 := `.A.`

        indices1 := findMatchStarts(pattern1, columns[i]) // M.S
        indices2 := findMatchStarts(pattern2, columns[i+1]) // .A.
        indices3 := findMatchStarts(pattern1, columns[i+2])// M.S
        fmt.Println(indices1,"-----")
        fmt.Println(indices2,"-----")
        fmt.Println(indices3,"-----")
        for _, match1 := range indices1 {
            for _, match2 := range indices2 {
                for _, match3 := range indices3 {
                    if match3 == match2 && match2 == match1 {
                        ans2++
                    }
                }
            }
        }
        fmt.Println(ans2)
    }

    
    fmt.Println(ans2)
}

//helper function to find match starting index only
func findMatchStarts(pattern, text string) []int {

    reg := regexp.MustCompile(pattern)
    start_idxs := []int{}
    for idx := 0; idx < len(text); idx++ {
        location := reg.FindStringIndex(text[idx:]) //find location at each slice

        if location == nil {
            break
        }

        absolute_idx := idx + location[0]
        flag := false
        for _, ele := range start_idxs {
            //fmt.Println(ele)
            if ele == absolute_idx {
                flag = true
            }
        }

        if !flag {
            start_idxs = append(start_idxs, absolute_idx)
        }

    }

    return start_idxs
}


//Helper function find XMAS
func find_xmas(input string) int {
    var total = 0

    //regex to find "XMAS"
    pattern := `XMAS`
    reg, _:= regexp.Compile(pattern)
    match_list := reg.FindAllString(input, -1) //-1: find all string -> return []string

    //fmt.Printf("%v\n", match_list)
    for _, match := range match_list {
        if match != "" {
            total++
        }
    }

    return total
}

//Helper function to reverse a string
func reverse (input string) string {

    //NOTE: there's no char in Go. Either use rune or byte
    output := []rune(input) // convert string into []rune (UTF-8) int32 (4bytes)
    for i := 0; i < len(output)/2; i++ {
        output[i], output[len(output) - 1 - i] = output[len(output) - 1 - i], output[i]

    }

    return string(output)
}
