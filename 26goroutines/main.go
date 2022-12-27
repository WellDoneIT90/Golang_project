package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // should be pointer
var mut sync.Mutex    // should be pointer

func main() {
	// go greeter("Hello")
	// greeter("World!")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
