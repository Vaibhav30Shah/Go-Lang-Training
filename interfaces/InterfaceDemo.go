package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printAreaAndPerimeter(shape shape) {
	fmt.Printf("Shape: %T\n", shape) //returns the type of the shape. format is package_name.struct_name. eg here interfaces.circle/rect
	fmt.Println("Area: ", shape.area())
	fmt.Println("Perimeter: ", shape.perimeter())
}

func InterfaceDemo() {
	c := circle{
		radius: 10,
	}

	r := rect{
		width:  10,
		height: 10,
	}

	printAreaAndPerimeter(c)
	printAreaAndPerimeter(r)

	var s shape
	s = circle{radius: 10}
	fmt.Printf("Type of s: %T\n", s) //types of interface s

	s = rect{
		width:  10,
		height: 10,
	}
	fmt.Printf("Type of s: %T\n", s) //type of s dynamically changes

	//embedded interface
	fmt.Println("#### EMBEDDED INTERFACE ####")
	EmbeddedInterface()
}
