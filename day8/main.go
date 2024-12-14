package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	//"errors"
)


type Matrix struct {
    position [][]string
}

func (m *Matrix) Init (rows [][]string) {
    m.position = make([][]string, len(rows))//Initialize a new position map
    
    for i, row := range rows {
        m.position[i] = make([]string, len(row))
        copy(m.position[i], row)
    }

}

func(m *Matrix) display() {
    for _, row := range m.position {
        fmt.Println(row)
    }
}
/**
func (m *Matrix) distance(pos1 [2]int, pos2 [2]int) (int, error) {
    var distance int = -1


    

    return distance, nil 
}
*/

func main() {

    data, err := os.ReadFile(os.Args[1])
    if err != nil{
        fmt.Println("Error reading file", os.Args[1])
        os.Exit(1)
    }
    fmt.Println(reflect.TypeOf(data))

    lines := strings.Split(strings.TrimSpace(string(data)), "\n") //Array of lines
    fmt.Println(reflect.TypeOf(lines))

    //Getting entries for the matrix
    var entries = make([][]string, len(lines))
    for i, line := range lines {
        entries[i] = strings.Split(line, "")//split to get each character
    }

    var matrix Matrix
    matrix.Init(entries)
    matrix.display()
}
