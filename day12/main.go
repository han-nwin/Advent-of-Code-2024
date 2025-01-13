package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Plant struct {
	kind     rune
	location [2]int
	neighbors int
}

// TODO:1. Implement get neighbor
// to get number of the neighbor (neighbor = plant of the same kind)
// we may not need this since we can just construct the neighbor when constructing []Plant (garden)
func (p *Plant) getNeighbors() {
	p.neighbors= 0
}

type Fence struct {
	plant       Plant
	numOfPlants int
}

func (f *Fence) area() int {
	return f.numOfPlants
}

// TODO: Calculate perimeter
func (f *Fence) perimeter() int {
	return 0
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("input file name not provided")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Cannot read input file")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println(lines)

	var plants []Plant

	for i, line := range lines {
		for j, k := range line {
			newPlant := Plant{
				kind:     k,
				location: [2]int{i, j},
			}
            // TODO: 1. Construct neighbors manually here
			newPlant.getNeighbors()

			plants = append(plants, newPlant)
		}
	}

	fmt.Println(plants)

}
