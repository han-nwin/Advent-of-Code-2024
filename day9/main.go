package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
    id string 
    size int
    startIdx int
    endIdx int
}

func main() {

    data, err := os.ReadFile(os.Args[1])
    if err != nil{
        fmt.Println("Error reading file", os.Args[1])
        os.Exit(1)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n") //Array of lines

    strArr := represent(lines[0])
   
    //Logic to move number to free space
    //Similar to quicksort median3 partition
    i := -1
    j := len(strArr)
    
    for {
        for {
            i++
            if strArr[i] == "." {
                break
            }
        }
        for {
            j--
            if strArr[j] != "." {
                break
            }
        }
        if i < j {
            strArr[i], strArr[j] = strArr[j], strArr[i] //swap
        } else {
            break
        }
    }

    ans1 := 0

    //Get result
    for i, char := range strArr {
        num, _ := strconv.Atoi(char)
        ans1 += (i * num)
    }
    fmt.Println("Answer 1:", ans1)

    //Part 2
    diskArr := represent(lines[0])

    var id string
    k := len(diskArr) - 1
    visited := make(map[string]bool)

    //Iterate through the disk backward
    for k >= 0 {
        //skip "."
        if diskArr[k] == "." {
            k--
            continue
        }
        
        //find the window of the file
        initial := k
        id = diskArr[k]

        for initial > 0 && diskArr[initial] == diskArr[initial -1] {
            initial--
        }
        window := k - initial + 1
        
        //Find free slot up to the initial
        idx := findFreeSlot(diskArr[:initial], window)
        
        //Swap
        if idx != -1 && idx < initial && !visited[id]{
            for offset := 0; offset < window; offset++ {
                diskArr[idx+offset] = id
            }
            // Clear original file location
            for j := initial; j <= k; j++ {
                diskArr[j] = "."
            }
            visited[id] = true
        }
        //Move backward to check next file
        k = initial - 1
    }

    //Get result
    ans2 := 0

    for s, char2 := range diskArr{
        num, _ := strconv.Atoi(char2)
        ans2 += (s * num)
    }
    fmt.Println("Answer 2:", ans2)


}

//Return the starting index of the free space
func findFreeSlot(diskArr []string, size int) int {
    if size <= 0 {
        return -1
    }

    freeCount := 0
    start := -1

    // Iterate through the array
    for i, char := range diskArr {
        if char == "." {
            if freeCount == 0 {
                start = i 
            }
            freeCount++
            if freeCount == size {
                return start //Return as soon as find a free space
            }
        } else {
            freeCount = 0 // Reset count if a non-free block is found
        }
    }

    return -1 // Can't find a space
}
func represent (line string) []string {
    var final []string

    fileID := 0
    for i, char := range line {
        size, _ := strconv.Atoi(string(char))
        var id string
        if i % 2 == 0 {
            id = strconv.Itoa(fileID)
            fileID++
        } else {
            id = "."
        }
        newFile := file{
            id:   id,
            size: size,
        }
        for range newFile.size {
            final = append(final, newFile.id)
        }
    }

    return final
}
