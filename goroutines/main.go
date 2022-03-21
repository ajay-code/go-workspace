package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	start := time.Now()
	fmt.Println("---start---")
	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
		"https://google.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)

	fmt.Println("---end---")
	elapsed := time.Since(start)
	log.Printf("it took %s", elapsed)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops failed endpoint.", endpoint)
	} else {
		fmt.Printf("%v status code for %v\n", res.StatusCode, endpoint)

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
	}

}
