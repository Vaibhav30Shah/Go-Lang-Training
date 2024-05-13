package concurrency

import "fmt"

func ChannelDemo() {
	var ch chan int
	fmt.Println(ch) //nil

	ch = make(chan int)
	fmt.Println(ch) //address in hex

	//channel is a 2 way messaging object.
	//It has 2 functions send and receive
	//Send:Transmits the value from 1 goroutine to other goroutine executing receive operation
	//<- channel operator

	//Send
	ch <- 10

	//Receive
	num := <-ch
	fmt.Println(<-ch)
	fmt.Println(num)

	close(ch) //closing a channel

	//channel always works with a goroutine. If a goroutine isnt stared it gives a fatal error of deadlock

}
