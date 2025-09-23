package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter a string\n> ")
		if !scanner.Scan() {
			break
		}
		input = scanner.Text()
		if input == "exit" || input == "q" {
			fmt.Printf("Bye:)\n")
			return
		}

		fmt.Printf("'%s' -> '%s'\n", input, backwards(input))
	}
}

func backwards(input string) string {
	tmp := []rune(input)
	rlen := len(tmp)
	for i := 0; i < rlen/2; i++ {
		tmp[i], tmp[rlen-1-i] = tmp[rlen-1-i], tmp[i]
	}

	return string(tmp)
}
