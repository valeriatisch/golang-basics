package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Ask the user to enter the file name
	fmt.Print("Enter the file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Read the file
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Convert the file content to a string and split it into words with strings.Fields()
	text := string(content)
	words := strings.Fields(text)

	// Use a map to store the the words and their number of occurrences
	wordCount := make(map[string]int)

	// Update word count map
	for _, word := range words {
		// Normalize the word (convert to lowercase)
		word = strings.ToLower(word)

		// Increment the word count
		wordCount[word]++
	}

	// Print the word count map
	fmt.Println("Word count:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
