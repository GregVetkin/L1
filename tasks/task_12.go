package main

import (
	"fmt"
)

func main() {
	var sl = []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]bool)
	set := make([]string, 0)

	for _, value := range sl {
		_, ok := m[value]
		if !ok {
			m[value] = true
			set = append(set, value)
		}
	}

	fmt.Println(set)
}