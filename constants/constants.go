package constants

import "fmt"

func ConstantDemo() {
	const x = 10 // no error in consts if also not used

	const (
		y = 20
		z
		s //automatic assigns the values of other constants as the upper constant's value
	)

	xx := 10
	yy := 0
	zz := xx / yy
	fmt.Println(zz) // no error at compile time, / by 0 error at runtime

	const aa = 10
	const bb = 0
	//const cc = aa / bb
	//fmt.Println(cc) // / by 0 error at compile time

	//const a int //error becoz we have to assign a value to a const
	//a=10 //error becoz we cant change the value of a constant

}
