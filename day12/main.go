package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Plant struct {
	kind      rune
	location  [2]int
	neighbors int
}

// Fence is just a collection of Plants
type Fence struct {
	plants []Plant
}

func (f *Fence) area() int {
	return len(f.plants)
}

func (f *Fence) perimeter() int {
	perimeter := 0
	for _, plant := range f.plants {
		// Calculate the perimeter contribution for each plant
		perimeter += 4 - plant.neighbors // Each neighbor reduces perimeter by 1
	}
	return perimeter
}

// TODO: Sides calculates the number of sides for the fence
func (f *Fence) sides() int {
    sides := 0
    return sides
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
	var fences []Fence

	for i, line := range lines {
		for j, k := range line {
			newPlant := Plant{
				kind:     k,
				location: [2]int{i, j},
			}
			//Construct neighbors manually here
			// Check neighboring cells in the grid directly
			if i > 0 && rune(lines[i-1][j]) == k { // Up
				newPlant.neighbors++
			}
			if i < len(lines)-1 && rune(lines[i+1][j]) == k { // Down
				newPlant.neighbors++
			}
			if j > 0 && rune(lines[i][j-1]) == k { // Left
				newPlant.neighbors++
			}
			if j < len(line)-1 && rune(lines[i][j+1]) == k { // Right
				newPlant.neighbors++
			}

			plants = append(plants, newPlant)
		}
	}

	fmt.Println(plants)

	//Constructing Fences
	// Create a 2D visited array
	visited := make([][]bool, len(lines))
	for i := range visited {
		visited[i] = make([]bool, len(lines[i]))
	}

	// NOTE: Helper function for flood-fill
	var floodFill func(i, j int, kind rune) Fence
	floodFill = func(i, j int, kind rune) Fence {
		stack := [][2]int{{i, j}}
        fence := Fence{}

		for len(stack) > 0 {
			// Pop a location from the stack
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			x, y := cur[0], cur[1]

			// Skip if out of bounds or already visited
			if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[x]) || visited[x][y] || rune(lines[x][y]) != kind {
				continue
			}

			// Mark as visited
			visited[x][y] = true

			// Add the corresponding plant
			for _, plant := range plants {
				if plant.location == [2]int{x, y} {
					fence.plants = append(fence.plants, plant)
					break
				}
			}

			// Add neighboring plants to stack (regardless of kind)
			stack = append(stack, [2]int{x - 1, y}) // Up
			stack = append(stack, [2]int{x + 1, y}) // Down
			stack = append(stack, [2]int{x, y - 1}) // Left
			stack = append(stack, [2]int{x, y + 1}) // Right
		}

		return fence
	}

	// Group fences by performing flood-fill
	for i, line := range lines {
		for j, char := range line {
			if !visited[i][j] {
				// Start a new fence for an unvisited plant
				fence := floodFill(i, j, rune(char))
				fences = append(fences, fence)
			}
		}
	}

	ans1 := 0
    ans2 := 0
	// Print the results
	fmt.Println("\nFences:")
	for i, fence := range fences {
		fmt.Printf("Fence %d: Area = %d, Perimeter = %d, Sides = %d\n", i+1, fence.area(), fence.perimeter(), fence.sides())
		ans1 += fence.area() * fence.perimeter()
        ans2 += fence.area() * fence.sides()
	}

	fmt.Println("ANSWER 1: ", ans1)
	fmt.Println("ANSWER 2: ", ans2)

}
