package main

import (
	"fmt"
	"unicode/utf8"
)


func Reverser(s string) string {
	symbolsNumber := utf8.RuneCountInString(s)

	r := make([]rune, symbolsNumber)
	symbolsNumber--

	for _, symbol := range s {
		r[symbolsNumber] = rune(symbol)
		symbolsNumber--
	}
	return string(r)
}


func main() {
	var somestring1 = "Hello Привет"
	var somestring2 = "главрыба"

	fmt.Println(Reverser(somestring1))
	fmt.Println(Reverser(somestring2))
}