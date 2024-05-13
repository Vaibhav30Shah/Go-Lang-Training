package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// mutexes is a solution for data race
func Mutexes() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	//1. declaring mutex
	var m sync.Mutex

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Lock()
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n--
			m.Unlock()
			wg.Done()
		}()
		//the order of go routines is unpredictable just like threads in java
	}
	wg.Wait()

	fmt.Println("n:", n)

}
