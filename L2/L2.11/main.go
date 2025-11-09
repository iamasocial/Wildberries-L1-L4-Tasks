package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "стол"}
	groups := findAnagramGroups(words)
	printGroups(groups)
}

func findAnagramGroups(words []string) map[string][]string {
	groups := make(map[string][]string)
	words = preprocess(words)
	for _, word := range words {
		key := sortWord(word)
		groups[key] = append(groups[key], word)
	}

	result := make(map[string][]string)
	for _, v := range groups {
		if len(v) < 2 {
			continue
		}

		sort.Strings(v)
		key := v[0]
		result[key] = v
	}

	return result
}

func sortWord(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func preprocess(words []string) []string {
	seen := make(map[string]struct{})
	for _, word := range words {
		low := strings.ToLower(word)
		seen[low] = struct{}{}
	}

	unique := make([]string, 0, len(seen))
	for w := range seen {
		unique = append(unique, w)
	}

	return unique
}

func printGroups(groups map[string][]string) {
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%q: %q\n", key, groups[key])
	}
}
