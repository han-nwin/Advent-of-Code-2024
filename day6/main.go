package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//write a step forward function depends on the shape
func main() {

    data, err := os.ReadFile(os.Args[1])
    if err != nil{
        fmt.Println("Error reading file", os.Args[1])
        os.Exit(1)
    }

    lines := strings.Split(string(data), "\n") //Array of lines
    lines = lines[:len(lines)-1] //Cut the last empty line
    
    ans1 := 0
    //BUILD the matrix
    //matrix
    var matrix [][]rune

    //guard
    var g guard
    g.Init() //initialize the guard
    g.Update() //Update the state

    fmt.Println("Position: ", g.curr_pos)
    fmt.Println("Ahead: ", g.ahead)
    //var barrier rune = '#'
    //var mark rune = 'X'
    
    //pattern to get position of the guard
    pattern := `v|\^|<|>`
    reg := regexp.MustCompile(pattern)
    var pos  = []int{-1,-1} // {row, col}

    //Iterate through every line
    for i, line := range lines {
        //Look for position of guard
        temp := reg.FindStringIndex(line)
        if temp != nil {
            pos[0] = i
            pos[1] = temp[0]
            g.shape = rune(line[temp[0]])
        }
        //Build the matrix
        char_list := []rune(line)
        matrix = append(matrix, char_list)
    }
    g.curr_pos = []int{pos[0], pos[1]}
    g.Update()
    //Print the matrix and starting pos
    fmt.Println("Position: ", g.curr_pos)
    /**
    fmt.Println("Ahead: ", g.ahead)
    fmt.Printf("Shape: %c\n", g.shape)
    for i, _ := range matrix {
        for _, r := range matrix[i]{
            fmt.Printf("%c ", r)
        }
        fmt.Println()
    }*/

    //RUNING LOGIC
    flag := false

    for !flag {
        // Before run mark current postion as X
        if matrix[g.curr_pos[0]][g.curr_pos[1]] != 'X' {
            matrix[g.curr_pos[0]][g.curr_pos[1]] = 'X'
        }

        //out at left or right
        if g.ahead[0] >= len(matrix) || g.ahead[0] < 0 {
            flag = true
            break
        }
        //out top or bottom
        if g.ahead[1] >= len(matrix[0]) || g.ahead[1] < 0 {
            flag = true
            break
        }
         
        //Check if there is barrier ahead
        if matrix[g.ahead[0]][g.ahead[1]] == '#' {
            g.turn_90()
        } else {
            g.go_forward(matrix)
        }

    }

    for i := range matrix {
        for _, char := range matrix[i] {
            if char == 'X' {
                ans1++
            }
        }

    } 
    fmt.Printf("Final Shape: %c\n", g.shape)
   /** for i := range matrix{
        for _, r := range matrix[i]{
            fmt.Printf("%c ", r)
        }
        fmt.Println()
    } */
    fmt.Println("Part 1: ", ans1)

    //Part 2
    //How to test if it's never escape?
    //NOTE: Visit the same position with the same direction

    matrix2 := copy_matrix(matrix)
    /**
    for i := range matrix2{
        for _, r := range matrix2[i]{
            fmt.Printf("%c ", r)
        }
        fmt.Println()
    }*/

    var ans2 = 0

    //guard 2
    var g2 guard
    g2.Init() //initialize the guard
    //Iterate through every line to initialize g2
    for i, line := range lines {
        //Look for position of guard
        temp := reg.FindStringIndex(line)
        if temp != nil {
            pos[0] = i
            pos[1] = temp[0]
            g2.shape = rune(line[temp[0]])
        }
    }
    initial_pos := []int{pos[0], pos[1]}
    g2.curr_pos = []int{pos[0], pos[1]}
    g2.Update()
    
    //Only put the obstacle in visited position in matrix of part 1 onto matrix 2
    //Test each visited position in matrix 1
    for i := range matrix2 {
        for j, char := range matrix2[i] {
            
            if char == 'X' {
                //don't put obstacle at initial position
                if initial_pos[0] == i && initial_pos[1] == j {
                    continue
                }

                //put a new obstacle in matrix2
               // fmt.Printf("Testing Obstruction at (%d, %d)\n", i, j)
                temp_matrix := copy_matrix(matrix2)
                temp_matrix[i][j] = '#'
                
                //Reset guard position before every loop
                g2.curr_pos = []int{pos[0],pos[1]} //reset position
                g2.shape = rune(lines[pos[0]][pos[1]]) //reset shape
                g2.Update() //reset ahead here

                visited := make(map[string]bool) // Tracks directions per position

                //Run
                for {
                //Capture current pos
                    position := fmt.Sprintf("%d,%d,%c", g2.curr_pos[0], g2.curr_pos[1], g2.shape)
                    
                    //cature loop
                    if visited[position] {
                        ans2++
                        break
                    }
                    visited[position] = true

                    //out at left or right
                    if g2.ahead[0] >= len(temp_matrix) || g2.ahead[0] < 0 {
                        break
                    }
                    //out top or bottom
                    if g2.ahead[1] >= len(temp_matrix[0]) || g2.ahead[1] < 0 {
                        break
                    }
                    
                    //Check if there is barrier ahead
                    if temp_matrix[g2.ahead[0]][g2.ahead[1]] == '#' {
                        g2.turn_90()
                    } else {
                        g2.go_forward(temp_matrix)
                    }
                    
                } //Done rune
            }
        }

    } // Done test whole matrix
    fmt.Println("Part 2:", ans2)

}

type guard struct {
    shape       rune
    curr_pos    []int
    ahead       []int
}

//init
func (g *guard) Init() {
    g.shape = '?'
    g.curr_pos = []int{-1,-1}
}
//update state
func (g *guard) Update () {
    switch g.shape {
    case 'v':
        g.ahead = []int{g.curr_pos[0] + 1, g.curr_pos[1]}
    case '<':
        g.ahead = []int{g.curr_pos[0], g.curr_pos[1] - 1}
    case '^':
        g.ahead = []int{g.curr_pos[0] - 1, g.curr_pos[1]}
    case '>':
        g.ahead = []int{g.curr_pos[0], g.curr_pos[1] + 1}

    }
}


//fucntion to determine if it's a guard type
func (g guard) is_guard() bool {
    if g.shape == 'v' || g.shape == '<' || g.shape == '>' || g.shape == '^' {
        return true
    } else {
        return false
    }
}

//function to turn 90 degree (just change the shape)
func (g *guard) turn_90() error {
    if !g.is_guard() {
        return errors.New("It's not a guard. Can't turn")
    }

    switch g.shape {
    case 'v':
        g.shape = '<'
    case '<':
        g.shape = '^'
    case '^':
        g.shape = '>'
    case '>':
        g.shape = 'v'

    }

    g.Update()
    return nil
}

//helper function to go forward in a matrix[][] 1 step, return err if out of bound
//flag = true if it goes outside of matrix
func (g *guard) go_forward(matrix [][]rune) (bool, error) {
    if g.shape == '?' {
        err := errors.New("Error going forward: wrong shape")
        return false, err
    }
    //Step forward
    switch g.shape {
    case 'v':
        g.curr_pos[0]++
    case '<':
        g.curr_pos[1]--
    case '^':
        g.curr_pos[0]--
    case '>':
        g.curr_pos[1]++
    }

    var flag = false
    
    //out at left or right
    if g.curr_pos[0] >= len(matrix) || g.curr_pos[0] < 0 {
        flag = true
    }
    //out top or bottom
    if g.curr_pos[1] >= len(matrix[0]) || g.curr_pos[1] < 0 {
        flag = true
    }
    g.Update()
    return flag, nil
}

//helper function to copy matrix
func copy_matrix (matrix [][]rune) [][]rune {
    copy_matrix := make([][]rune, len(matrix))
    for i := range matrix {
        copy_matrix[i] = make([]rune, len(matrix[i]))
        copy(copy_matrix[i], matrix[i])
    }
    return copy_matrix
}

//helper function to mark position
func markPosition(matrix [][]rune, pos []int, direction rune) {
    switch direction {
    case '^', 'v': // Vertical movement
        if matrix[pos[0]][pos[1]] == '-' { // Already marked as horizontal
            matrix[pos[0]][pos[1]] = '+'
        } else {
            matrix[pos[0]][pos[1]] = '|'
        }
    case '<', '>': // Horizontal movement
        if matrix[pos[0]][pos[1]] == '|' { // Already marked as vertical
            matrix[pos[0]][pos[1]] = '+'
        } else {
            matrix[pos[0]][pos[1]] = '-'
        }
    }
}
