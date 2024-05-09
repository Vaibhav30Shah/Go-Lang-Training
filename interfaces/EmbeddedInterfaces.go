package interfaces

import (
	"fmt"
	"math"
)

type twoD interface {
	area2D() float64
}

type threeD interface {
	volume3D() float64
}

type object interface {
	twoD
	threeD
	getColor() string
}

type cube struct {
	side  float64
	color string
}

func (c cube) area2D() float64 {
	return 6 * (c.side * c.side)
}

func (c cube) volume3D() float64 {
	return math.Pow(c.side, 3)
}

func (c cube) getColor() string {
	return c.color
}

func measure(o object) (float64, float64) {
	a := o.area2D()
	v := o.volume3D()
	return a, v
}

func EmbeddedInterface() {
	c := cube{side: 2., color: "red"}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volume: %v\n", a, v)
}
