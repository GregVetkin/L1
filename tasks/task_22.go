package main

import (
	"fmt"
	"math/big"
)

func main() {
	// для работы с большии числами можно использовать встроенный пакет big

	x := new(big.Int)
	x.SetString("4444444444444444444444444444444444444444444444444444444444444444444", 10)

	y := new(big.Int)
	y.SetString("2222222222222222222222222222222222222222222222222222222222222222222", 10)

	result := new(big.Int)


	// сложение
	result.Add(x, y)
	fmt.Println(result)

	// вычитание
	result.Sub(y, x)
	fmt.Println(result)

	// умножение
	result.Mul(y, x)
	fmt.Println(result)

	// деление
	result.Div(x, y)
	fmt.Println(result)


}