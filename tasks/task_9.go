package main

import "fmt"


func readChanel(numbers []int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, number := range numbers {
			outCh <- number
		}
	}()
	return outCh
}

func squareChanel(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for number := range inCh {
			outCh <- number * number
		}
	}()
	return outCh
}


func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	in := readChanel(numbers)
	out := squareChanel(in)

	for number := range out {
		fmt.Println(number)
	}
}
