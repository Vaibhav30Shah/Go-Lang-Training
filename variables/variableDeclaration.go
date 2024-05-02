package variables

import (
	"fmt"
)

const abshd = 90

func VariableDeclaration() {
	var x int = 10 //using var initialization of variable

	var xx = "Ab10" //without using data type

	var s string //declaration
	s = "Hello"  //initialization

	println(x)     //prints the output in std err
	fmt.Println(s) //prints the output in std out
	fmt.Printf(xx)

	y := true //shorthand variable initialization
	println(y)

	_ = "No error" //using _ we will not get error for non used vars

	//Multiple declarations
	a, b := "Vaibhav", 194
	println(a, b)

	var (
		name   string
		age    int
		salary float64
		gender bool
		aaa    rune
	)
	fmt.Println(name, age, salary, gender, aaa)
}
