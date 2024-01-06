// Get content type of sites (using channels)
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s -> error: %s\n", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s", url, ctype)
}

func sender(ch chan<- string, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- value
	fmt.Printf("Sent: %d\n", value)
}

func receiver(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-ch
	fmt.Printf("Received: %d\n", val)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url)
	}

	// TODO: Wait using channel
	var wg sync.WaitGroup

	// Add a sender
	for _, url := range urls {
		wg.Add(1)
		go sender(ch, url, &wg)
	}

	// Add three receivers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go receiver(ch, &wg)
	}

	// Wait for all goroutines to finish

	wg.Wait()
}
