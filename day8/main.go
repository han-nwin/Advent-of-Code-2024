package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)


type Matrix struct {
    position [][]string
    numRow int
    numCol int
}

func (m *Matrix) Init (rows [][]string) {
    m.position = make([][]string, len(rows))//Initialize a new position map
    m.numRow = len(rows)
    for i, row := range rows {
        m.numCol = len(row)
        m.position[i] = make([]string, len(row))
        copy(m.position[i], row)
    }

}

func(m *Matrix) display() {
    fmt.Println(m.numRow, m.numCol)
    for _, row := range m.position {
        fmt.Println(row)
    }
}

func (m *Matrix) distance(pos1 [2]int, pos2 [2]int) [2]int {
    vertical := pos1[0] - pos2[0]
    horizontal := pos1[1] - pos2[1]
    distance := [2]int{vertical, horizontal}

    return distance
}

func (m *Matrix) antiNodePositions (pos1 [2]int, pos2 [2]int) [2][2]int {
    distance := m.distance(pos1,pos2)
    
    var anti1Pos [2]int
    var anti2Pos [2]int
    var antiPos [2][2]int

    anti1Pos[0] = pos1[0] + distance[0]
    anti1Pos[1] = pos1[1] + distance[1]


    anti2Pos[0] = pos2[0] - distance[0]
    anti2Pos[1] = pos2[1] - distance[1]

    antiPos[0] = anti1Pos
    antiPos[1] = anti2Pos

    return antiPos
}

func (m *Matrix) addAntiNodes (pos1 [2]int, pos2 [2]int) {
    antiPos := m.antiNodePositions(pos1, pos2)
    for _, node := range antiPos {
        if node[0] >=0 && node[0] < m.numRow && node[1] >=0 && node[1] < m.numCol {
            m.position[node[0]][node[1]] = "#"
        }
    }
}

//find all match and return positions array of the matches
func (m *Matrix) findAllMatch (shape string) [][2]int {
    var matchList [][2]int

    for i, row := range m.position {
        for j, value := range row {
            if shape == value {
                matchList = append(matchList, [2]int{i, j})
            }
        }
    }
    return matchList 
}

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

    var antennaMap = make(map[string][][2]int)

    var visited = make(map[string]bool)
    for _, row := range matrix.position {
        for _, shape := range row {
            if shape != "." && shape != "#" {
                if !visited[shape] {
                    antennaMap[shape] = matrix.findAllMatch(shape)
                    visited[shape] = true
                }
            }
        }
    }
    fmt.Println(antennaMap)
    
    ans1 := 0

    for _, array := range antennaMap {
        for i :=0; i < len(array); i++ {
            for j := i+1; j < len(array); j++ {
                matrix.addAntiNodes(array[i], array[j])
            }

        }
    }

    matrix.display() 
    for _, row := range matrix.position {
        for _, shape := range row {
            if shape == "#"{
                ans1++
            }
        }
    }
    fmt.Println("Part 1: ",ans1)
}
