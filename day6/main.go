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
    var pos []int = []int{-1,-1} // {row, col}

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
    g.curr_pos = pos
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
    var e error

    for !flag {
        //Before run mark the current pos as 'X'
        if matrix[g.curr_pos[0]][g.curr_pos[1]] != 'X'{
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
        if matrix[g.ahead[0]][g.ahead[1]] == '#' && !flag {
            g.turn_90()
        }

        flag, e = g.go_forward(matrix)
        
        if e != nil {
            fmt.Println(e)
            os.Exit(1)
        }
        fmt.Println(flag,e)
        fmt.Println("New position: ", g.curr_pos)

        if !flag {
            matrix[g.curr_pos[0]][g.curr_pos[1]] = g.shape
        }

        /**
        fmt.Printf("Shape: %c\n", g.shape)
        for i, _ := range matrix {
            for _, r := range matrix[i]{
                fmt.Printf("%c ", r)
            }
            fmt.Println()
        }
        */

    }
    for i, _ := range matrix {
        for _, char := range matrix[i] {
            if char == 'X' {
                ans1++
            }
        }

    } 
    fmt.Println("Part 1: ", ans1)
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

    flag := false
    
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
