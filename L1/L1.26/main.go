package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"abcd", "abCdefAaf", "aabcd", "Hello", "He=lo"}

	for _, word := range words {
		fmt.Printf("%q -> %v\n", word, CheckUnique(word))
	}
}

func CheckUnique(str string) bool {
	str = strings.ToLower(str)
	seen := make(map[rune]struct{})

	for _, r := range str {
		if _, exists := seen[r]; exists {
			return false
		}

		seen[r] = struct{}{}
	}

	return true
}
