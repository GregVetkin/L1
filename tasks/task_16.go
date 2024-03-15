package main

import "fmt"

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := 0
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for indx, _ := range arr {
		if arr[indx] < arr[right] {
			arr[left], arr[indx] = arr[indx], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	Quicksort(arr[:left])
	Quicksort(arr[left+1:])

	return arr
}


func main() {
	a := []int{0, 8, -5, 100, 12, -33, 16, -1000}
	Quicksort(a)
	fmt.Println(a) // -1000 -33 -5 0 8 12 16 100
}