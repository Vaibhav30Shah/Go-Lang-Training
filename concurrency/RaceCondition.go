package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RaceCondition() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
		//the order of go routines is unpredictable just like threads in java
	}
	wg.Wait()

	fmt.Println("n:", n)
}
