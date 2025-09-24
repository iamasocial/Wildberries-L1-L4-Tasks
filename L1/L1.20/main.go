package main

import "fmt"

func main() {
	str := "sun dog snow"
	r := []rune(str)
	fmt.Printf("Original str: %q\n", string(r))
	reverseWords(r, ' ')
	fmt.Printf("Reversed str: %q\n", string(r))
}

func reverseRunes(r []rune) {
	rlen := len(r)
	for i, j := 0, rlen-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func reverseWords(r []rune, sep rune) {
	reverseRunes(r)
	start := 0
	rlen := len(r)

	for i := 0; i <= rlen; i++ {
		if i == rlen || r[i] == sep {
			reverseRunes(r[start:i])
			start = i + 1
		}
	}
}
