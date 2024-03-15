package main

import (
	"fmt"
	"strings"
)

func UniqueString(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]struct{})

	for _, char := range s {
		_, ok := m[char] 
		if ok {
			return false
		}
		m[char] = struct{}{}
	}

	return true
}




func main() {
	var strings = []string{
		"abcd",			//	true
		"abCdefAaf",	// false
		"aabcd",		// false
		"見てくれてありがとう",	// false
	}

	for _, s := range strings {
		fmt.Println(s, "-",UniqueString(s))
	}
}