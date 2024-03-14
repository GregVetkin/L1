package main

import (
	"fmt"
	"reflect"
)


func varTypeFMT(x interface{}) {
	fmt.Printf("fmt: %T\n", x)
}


func varTypeSWITCH(x interface{}) {
	fmt.Print("switch: ")

	switch t := x.(type) {
	case int:
		fmt.Println("int | value:", t)
	case string:
		fmt.Println("string | value:", t)
	case bool:
		fmt.Println("bool | value:", t)
	case chan any:
		fmt.Println("chan interface{} | value:", t)
	default:
		fmt.Println("unknown | value:", t)
	}
}

func varTypeREFLECT(x interface{}) {
	xType := reflect.TypeOf(x)
	fmt.Printf("reflect: %s", xType)
}




func main() {
	var a interface{} = 1
	var b interface{} = "hi"
	var c interface{} = true
	var d interface{} = make(chan any)

	data := []any{a, b, c, d}

	for _, v := range data {
		varTypeFMT(v)
		varTypeSWITCH(v)
		varTypeREFLECT(v)
		fmt.Println()
	}


}