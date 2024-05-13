package concurrency

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s is Down\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
	} else {
		defer resp.Body.Close() //used defer because we want to close just before main closes
		s := fmt.Sprintf("%s is Up. Status code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] //separates the website url by //
			file += ".txt"
			s += fmt.Sprintf("Writing resp body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				//log.Fatal(err)
				s += "Error writing to file"
				c <- s
			}
			s += fmt.Sprintf("%s is UP.\n", url)
			c <- s
		}
	}
	wg.Done()
}

func MainUrlChecker() {
	urls := []string{"https://www.google.com", "https://www.golang.org", "https://facebook.com"}

	c := make(chan string)
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg, c)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of goroutines: ", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
	wg.Wait()
}
