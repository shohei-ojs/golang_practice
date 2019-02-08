package main

import (
	"fmt"
)

// How to missing number of 1 to 10
func main() {
	// missing number to five
	numbers := [10]int{1,2,3,4,6,7,8,9,10}
	n := len(numbers)
	expect := n*(n+1)/2
	fmt.Println(expect - sum(numbers))

}

func sum (arr [10]int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}