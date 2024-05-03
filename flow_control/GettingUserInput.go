package flow_control

import (
	"fmt"
	"os"
	"strconv"
)

func GettingUserInput() {
	println(os.Args[0])
	fmt.Println("Enter your name:", os.Args[1])
	val, err := strconv.Atoi("222")
	if err != nil {
		fmt.Println("Invalid input")
	}
	fmt.Println("Answer", val)

	val2 := strconv.Itoa('a')
	fmt.Println("Answer", val2)
}
