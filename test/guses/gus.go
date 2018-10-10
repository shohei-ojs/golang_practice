package guses

import (
	"fmt"
)

type Gus interface {
	Name()       string
	Combustion() string
}

type Unleaded struct {
	Name    string
	Quality float32
}

func NewUnleaded(n string, q float32) *Unleaded {
	return &Unleaded{n,q}
}

func (u *Unleaded) Combustion() string {
	return fmt.Sprintf("%sのQualityは%s", u.Name, u.Quality)
}