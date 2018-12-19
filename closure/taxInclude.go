package main

import (
	"flag"
	"fmt"
)

var (
	price = flag.Int("price", 1000, "value price")
)

func taxIncluder(tax float64) func(price int) int {
	t := tax
	return func(price int) int {
		return int(float64(price) * t)
	}
}

func main() {
	flag.Parse()

	taxop := taxIncluder(1.08)
	fmt.Println(taxop(*price))

	taxop2 := taxIncluder(1.235)
	fmt.Println(taxop2(*price))
}
