package main

import "fmt"

func BinarySearch(arr []int, elem int) (int, bool) {
	start := 0
	end := len(arr) - 1

	for start <= end {
		middle := (start + end) / 2
		if arr[middle] == elem {
			return middle, true
		}
		if arr[middle] < elem {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return 0, false
}

func main() {
	var arr = []int{-1, 0, 1}
	fmt.Println(BinarySearch(arr, -1))	// 0 true
	fmt.Println(BinarySearch(arr, 0))	// 1 true
	fmt.Println(BinarySearch(arr, 1))	// 2 true
	fmt.Println(BinarySearch(arr, 8))	// 0 false
}