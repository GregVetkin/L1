package main

import (
	"fmt"
	"time"
)

func Sleep1(t time.Duration) {
	timer := time.NewTimer(t)
	<- timer.C
}

func Sleep2(t time.Duration) {
	<- time.After(t)
}

func main() {
	duration := 500 * time.Millisecond

	start1 := time.Now()
	Sleep1(duration)
	sl1 := time.Since(start1)

	start2 := time.Now()
	Sleep2(duration)
	sl2 := time.Since(start2)

	start3 := time.Now()
	time.Sleep(duration)
	sl3 := time.Since(start3)


	fmt.Printf("Sleep1: %v - have to: %v\n", sl1, duration)
	fmt.Printf("Sleep2: %v - have to: %v\n", sl2, duration)
	fmt.Printf("time.Sleep: %v - have to: %v\n", sl3, duration)
}