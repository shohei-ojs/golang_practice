package main

import(
	"os"
	"fmt"
)

func main() {
	render, err := os.Open("../image/gopher.png")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(*render)
}