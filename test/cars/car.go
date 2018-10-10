package cars

import (
	"fmt"
)

type BigCar struct {
	Name string
	Size size
}

type size struct {
	height, width, length int
}

func NewBigCar(name string, size size) *BigCar{
	return &BigCar{name,size}
}

func (bc *BigCar) WathYourName() string{
	return fmt.Sprintf("%s says quack", bc.Name)
}