package main

import (
	"fmt"
	"./bigint"
)

var p = fmt.Println

func main() {
	a := bigint.FromString("3635634663456345635634565363644356345748745673457")
	b := bigint.FromString("2453609867590753475927458790347502743058274034563456431")

	p(a.Mul(b))
}