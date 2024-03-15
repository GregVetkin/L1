package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}

	fmt.Printf("Slice: %v; lenght: %d, capacity: %d\n", slice, len(slice), cap(slice))

	i := 1	// index 1 value 2
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Printf("Slice: %v; lenght: %d, capacity: %d\n", slice, len(slice), cap(slice))
}