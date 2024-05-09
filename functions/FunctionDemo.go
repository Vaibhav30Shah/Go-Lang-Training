package functions

import "fmt"

func FunctionDemo() {
	var a float64
	_, a = f1(10, 20)
	fmt.Println(a)

	variadicFunctionDemo(10., 11., 12.2222)

}

func f1(a int, b int) (s int, z float64) {
	s = a + b
	return
}

// ... means a slice
func variadicFunctionDemo(val ...float64) {
	sum := 0.
	prod := 1.
	for _, v := range val {
		sum += v
		prod *= v
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", prod)
}
