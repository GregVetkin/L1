package main

import (
	"errors"
	"fmt"
)

func ChangeBit(value, i int) (int, error) {
	if i < 1 || i > 65 {
		return value, errors.New("Changed bit has to be between 1 and 64")
	}
	i -= 1

	value ^= (1 << i)
	return value, nil
}



func main() {
	numberBefore := 7
	fmt.Printf("%d -> %b\n", numberBefore, numberBefore)
	numberAfter, err := ChangeBit(numberBefore, 5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d -> %b\n", numberAfter, numberAfter)

	numberAfter2, err := ChangeBit(numberAfter, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d -> %b\n", numberAfter2, numberAfter2)
}