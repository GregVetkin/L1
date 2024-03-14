package main

import (
	"fmt"
)

func main () {
	var temps = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempgroups := make(map[int][]float64)

	for _, temp := range temps {
		tempstep := int(temp / 10) * 10
		tempgroups[tempstep] = append(tempgroups[tempstep], temp)
	}

	fmt.Println(tempgroups[-20])
	fmt.Println(tempgroups[10])
	fmt.Println(tempgroups[20])
}