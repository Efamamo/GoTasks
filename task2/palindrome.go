package task2

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	word := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")

	i, j := 0, len(word)-1
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i, j = i+1, j-1
	}

	return true

}
