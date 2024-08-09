package task2

import (
	"regexp"
	"strings"
)

func WordFrequency(input string) map[string]int {
	// Convert the string to lowercase
	input = strings.ToLower(input)

	// Remove punctuation using regex
	cleanedInput := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(input, " ")

	// Split the string into words
	words := strings.Fields(cleanedInput)

	// Create a map to store word frequencies
	frequencyMap := map[string]int{}

	// Count the frequency of each word
	for _, word := range words {
		frequencyMap[word]++
	}

	return frequencyMap
}
