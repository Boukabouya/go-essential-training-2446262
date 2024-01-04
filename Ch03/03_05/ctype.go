// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

	// Continue with the rest of your code, knowing that the HTTP request and response body reading were successful.

}

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url

// contentType retrieves the Content-Type header from an HTTP HEAD request to the specified URL.

func contentType(url string) (string, error) {
	// Make an HTTP HEAD request
	response, err := http.Head(url)
	if err != nil {
		return "", fmt.Errorf("HTTP HEAD request failed: %v", err)
	}

	// Check if the request was successful (status code 200)
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP HEAD request failed with status code: %d", response.StatusCode)
	}

	// Retrieve the Content-Type header
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("Content-Type header not found in the response")
	}

	return contentType, nil
}
