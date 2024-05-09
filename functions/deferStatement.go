package functions

import "fmt"

func func1() {
	fmt.Println("This is f1")
}

func func2() {
	fmt.Println("This is f2")
}

func func3() {
	fmt.Println("This is f3")
}

func DeferStatement() {
	defer func1() //this will be called last
	func2()

	fmt.Println("This is defer")

	defer func3() //this will be called second last

	//in defer the calls are stored in a stack and follow lifo. so it executes in reverse order.
}
