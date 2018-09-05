package builder

import (
	"fmt"
)

// 产品是Car
type Car struct {
	Type  string
	Brand string
	Color string
}

func (c Car) Move() {
	fmt.Println("moving moving", c.Brand, c.Type, c.Color)
}

type Builder interface {
	SetType(t string) Builder
	AddBrand(b string) Builder
	PaintColor(c string) Builder
	Build() Car
}

type CarBuilder struct {
	Acar Car
}

func (cb CarBuilder) SetType(t string) Builder {
	cb.Acar.Type = t
	return cb
}

func (cb CarBuilder) AddBrand(b string) Builder {
	cb.Acar.Brand = b
	return cb
}

func (cb CarBuilder) PaintColor(c string) Builder {
	cb.Acar.Color = c
	return cb
}

func (cb CarBuilder) Build() Car {
	return cb.Acar
}

type Director struct {
	Builder Builder
}
