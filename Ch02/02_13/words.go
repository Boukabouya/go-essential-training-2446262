package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	textSlice := strings.Fields(text) // Split text into words
	wordFrequency := make(map[string]int)

	// Convert each word to lowercase and count frequency
	for _, word := range textSlice {
		lowercaseWord := strings.ToLower(word)
		wordFrequency[lowercaseWord]++
	}

	// Display the word frequencies
	for word, count := range wordFrequency {
		fmt.Printf("The word %s appears %d times\n", word, count)
	}
}
