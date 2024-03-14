package main

import (
	"fmt"
)

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	set2 := []int{1, 2, 11, 12, 13, 14, 15, 16}

	s := make(map[int]int)
	intersection := make([]int, 0)

	for _, n := range set1 {
		s[n] += 1
	}
	for _, n := range set2 {
		s[n] += 1
	}
	for key, value := range s {
		if value > 1 {
			intersection = append(intersection, key)
		}
	}
	fmt.Println(intersection)
}