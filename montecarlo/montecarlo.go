package main

import (
	"fmt"
)

// Condition ...
type Condition string

const (
	first Condition = "first"
	won   Condition = "won"
	lose  Condition = "lose"
)

var x []int = []int{1, 2, 3}

// Get ...
func Get(c Condition) int {
	switch c {
	case won:
		return func() int {
			x = x[1 : len(x)-1]
			if len(x) <= 1 {
				x = []int{1, 2, 3}
			}
			return x[0] + x[len(x)-1]
		}()
	case lose:
		return func() int {
			x = append(x, x[0]+x[len(x)-1])
			return x[0] + x[len(x)-1]
		}()
	default:
		return 4
	}
}

func main() {
	f := fmt.Println
	f(Get(first))
	f(Get(lose))
	f(Get(lose))
	f(Get(won))
	f(Get(lose))
	f(Get(lose))
	f(Get(won))

}
