package variables

import "fmt"

func TypesOfVariables() {
	var x = 5
	var y = 12.5
	var z = int(y)
	var p = byte(y)

	var hi_how_are_you = 42 //output will come but this is not idiomatic(as per go standard). Use hiHowAreYou instead. Dont use _ in between name

	var b float64
	fmt.Println(b == 0.0)

	fmt.Printf("Hi my name is %s and my age is %+d. My CPI are %+.2f\n", "Vaibhav", -21, 9.89)
	fmt.Printf("I am from %q\n", "Bhuj")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(p)
	fmt.Println(hi_how_are_you)
}
