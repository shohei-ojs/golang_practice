package main

import (
	"fmt"
)

func a(pi float64) func(radius float64) float64 {
	p := pi
	return func(radius float64) float64 {
		return p * radius * 2
	}
}

func main() {
	calculator := a(3.14)
	result := calculator(3)
	fmt.Printf("size of area : %v\n", result)

	calculator2 := a(3)
	result = calculator2(3)
	fmt.Printf("size of area : %v\n", result)
}
