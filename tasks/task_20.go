package main

import (
	"fmt"
	"strings"
)

func ReverserWords(s string) string {
	var result strings.Builder

	words := strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		if i != 0 {
			result.WriteString(" ")
		}
	}
	return result.String()
}


func main() {
	var somestring1 = "snow dog sun"
	fmt.Println(ReverserWords(somestring1))

	var somestring2 = "Hello friend"
	fmt.Println(ReverserWords(somestring2))
}