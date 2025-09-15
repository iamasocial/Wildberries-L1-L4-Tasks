package main

import "fmt"

func main() {
	words := []string{"golang", "python", "java", "golang", "c++", "python"}
	set := NewSet(words)

	for word := range set {
		fmt.Println(word)
	}
}

func NewSet(words []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, word := range words {
		set[word] = struct{}{}
	}

	return set
}
