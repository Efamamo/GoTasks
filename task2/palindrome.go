package task2

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^\w]`)
	word := re.ReplaceAllString(s, "")

	i, j := 0, len(word)-1
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i, j = i+1, j-1
	}

	return true

}
