package main

import (
    "fmt"
    "container/list" //doubly linked list

)


func main() {

    queue := list.New()


	// Initialize a new list
	l := list.New()

	// Add elements of different types
	l.PushBack(10)         // int
	l.PushBack("hello")    // string
	l.PushBack(3.14)       // float64
	l.PushBack(true)       // bool

	// Print elements and their types
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("Value: %v, Type: %T\n", e.Value, e.Value)
	}

	// Access and remove elements
	front := l.Front() // Get the first element
	fmt.Println("Front element:", front.Value)

	l.Remove(l.Front()) // Remove the front element

	fmt.Println("After removing the front element:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}


}
