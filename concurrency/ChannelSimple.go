package concurrency

import "fmt"

func ChannelSimple() {
	//bidirectonal channel
	c := make(chan int)

	//unidirectional channel
	c1 := make(<-chan int) //only recv
	c2 := make(chan<- int) //only send

	fmt.Printf("%T, %T, %T\n", c, c2, c1)

	go func(n int, ch chan int) {
		ch <- n
	}(10, c)

	n := <-c
	fmt.Printf("Value recd: %v\n", n)
	defer fmt.Printf("Main is exiting")
}
