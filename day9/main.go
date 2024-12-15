
package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {

    data, err := os.ReadFile(os.Args[1])
    if err != nil{
        fmt.Println("Error reading file", os.Args[1])
        os.Exit(1)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n") //Array of lines


}
