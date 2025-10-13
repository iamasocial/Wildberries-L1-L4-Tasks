package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"
)

func main() {
	strs := []string{`a4bc2d5e`, `abcd`, `45`, ``, `qwe\4\5`, `qwe\45`, `qwe45`}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for _, str := range strs {
		unpacked, err := unpack(str)

		if err != nil {
			fmt.Fprintf(w, "failed to unpack\t%q\t%v\n", str, err)
			continue
		}

		fmt.Fprintf(w, "packed string\t%q\tunpacked string\t%q\n", str, unpacked)
	}

	w.Flush()
	fmt.Println(strings.Repeat("a", 5))
}

func unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	var prev rune
	var hasPrev, escaped bool

	if s[len(s)-1] == '\\' {
		return "", fmt.Errorf("invalid string:\tends with escape character")
	}

	for _, r := range s {
		if r == '\\' {
			if escaped {
				return "", fmt.Errorf("invalid string:\tconsecutive escape characters")
			}
			escaped = true
			continue
		}

		if escaped {
			result.WriteRune(r)
			prev = r
			hasPrev = true
			escaped = false
			continue
		}

		if unicode.IsDigit(r) {
			if !hasPrev {
				return "", fmt.Errorf("invalid string:\tno symbol before digit")
			}

			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}

			result.WriteString(strings.Repeat(string(prev), count-1))

			hasPrev = false

			continue
		}

		result.WriteRune(r)
		prev = r
		hasPrev = true
	}

	return result.String(), nil
}
