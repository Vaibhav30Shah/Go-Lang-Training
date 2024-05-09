package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doSomething1() {
	fmt.Println("1st do something")
}

func doSomething2(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		println(i)
		time.Sleep(time.Second)
	}
	wg.Done()
	println("Exiting goroutine")
}

func GoRoutineDemo() {
	var wg sync.WaitGroup
	wg.Add(1)
	//normal function invocation
	doSomething1()

	//spawning a goroutine
	go doSomething2(&wg)

	fmt.Println("No of CPUS:", runtime.NumCPU())
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	wg.Wait()
	println("Exiting main")
}
