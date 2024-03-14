package main

import (
	"fmt"
)


func main() {
	a := 100
	b := 200

	fmt.Printf("A=%d, B=%d\n", a, b)

	a, b = b, a

	fmt.Printf("A=%d, B=%d\n", a, b)
}