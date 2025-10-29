package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var monthOrder = map[string]int{
	"jan": 1, "feb": 2, "mar": 3, "apr": 4,
	"may": 5, "jun": 6, "jul": 7, "aug": 8,
	"sep": 9, "oct": 10, "nov": 11, "dec": 12,
}

var humanReadableSuffixes = map[string]int{
	"K": 1024, "M": 1024 * 1024,
	"G": 1024 * 1024 * 1024, "T": 1024 * 1024 * 1024 * 1024,
}

func main() {
	opts := Options{}
	pflag.IntVarP(&opts.Column, "column", "k", 0, "sort by column number (1-based, tab separated)")
	pflag.BoolVarP(&opts.Numeric, "numeric", "n", false, "sort numerically")
	pflag.BoolVarP(&opts.Reverse, "reverse", "r", false, "reverse sort order")
	pflag.BoolVarP(&opts.Unique, "unique", "u", false, "include only unique lines")
	pflag.BoolVarP(&opts.MonthSort, "month", "m", false, "sort by month name")
	pflag.BoolVarP(&opts.IgnoreTrailing, "blanks", "b", false, "ignore trailing blanks")
	pflag.BoolVarP(&opts.CheckSorted, "sorted", "c", false, "check if input is sorted")
	pflag.BoolVarP(&opts.HumanNumeric, "human", "h", false, "human-readable numeric sort (2K, 4M, etc)")
	pflag.Parse()

	var reader io.Reader = os.Stdin
	if args := pflag.Args(); len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading:", err)
		os.Exit(1)
	}

	lines = preProcessLines(lines, opts.Column, opts.IgnoreTrailing)
	sortable := SortableLines{lines: lines, options: opts}

	if opts.CheckSorted {
		if checkSorted(sortable) {
			printLines(sortable.lines)
			return
		}
		fmt.Fprintf(os.Stderr, "Input not sorted, sorting now\n")
	}
	sort.Stable(sortable)

	lines = postProcessLines(lines, opts.Unique)

	printLines(lines)
}

// Options defines sorting options such as column, numeric, etc.
type Options struct {
	Column         int
	Numeric        bool
	Reverse        bool
	Unique         bool
	MonthSort      bool
	IgnoreTrailing bool
	HumanNumeric   bool
	CheckSorted    bool
}

// SortableLines represents a sortalbe collection of line with options
type SortableLines struct {
	lines   []string
	options Options
}

func (s SortableLines) Len() int {
	return len(s.lines)
}

func (s SortableLines) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func (s SortableLines) Less(i, j int) bool {
	a := s.lines[i]
	b := s.lines[j]

	cmp := 0

	aCol := getColumn(a, s.options.Column)
	bCol := getColumn(b, s.options.Column)

	switch {
	case s.options.Numeric:
		cmp = compareNumeric(aCol, bCol)
	case s.options.MonthSort:
		cmp = compareMonth(aCol, bCol)
	case s.options.HumanNumeric:
		cmp = compareHumanNumeric(aCol, bCol)
	default:
		cmp = strings.Compare(aCol, bCol)
	}

	if s.options.Reverse {
		return cmp > 0
	}

	return cmp < 0
}

func compareNumeric(a, b string) int {
	af, err1 := strconv.ParseFloat(a, 64)
	bf, err2 := strconv.ParseFloat(b, 64)

	if err1 != nil || err2 != nil {
		return strings.Compare(a, b)
	}

	if af < bf {
		return -1
	}

	if af == bf {
		return 0
	}

	return 1
}

func getColumn(line string, col int) string {
	if col <= 0 {
		return line
	}

	fields := strings.Split(line, "\t")

	if col <= len(fields) {
		field := fields[col-1]
		return field
	}

	return ""
}

func compareMonth(a, b string) int {
	am, ok1 := monthOrder[strings.ToLower(a)]
	bm, ok2 := monthOrder[strings.ToLower(b)]

	if !ok1 || !ok2 {
		return strings.Compare(a, b)
	}

	if am < bm {
		return -1
	}

	if am == bm {
		return 0
	}

	return 1
}

func parseHuman(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string)")
	}

	mult := 1
	suffix := s[len(s)-1:]

	if v, ok := humanReadableSuffixes[suffix]; ok {
		s = s[:len(s)-1]
		mult = v
	}

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return int(value * float64(mult)), nil
}

func compareHumanNumeric(a, b string) int {
	ah, err1 := parseHuman(a)
	bh, err2 := parseHuman(b)
	if err1 != nil || err2 != nil {
		return strings.Compare(a, b)
	}

	if ah < bh {
		return -1
	}

	if ah == bh {
		return 0
	}

	return 1
}

func checkSorted(s SortableLines) bool {
	for i := 0; i < s.Len()-1; i++ {
		if s.Less(i+1, i) {
			return false
		}
	}

	return true
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}

func preProcessLines(lines []string, col int, ignoreTrailing bool) []string {
	if ignoreTrailing {
		lines = trimRightSpaces(lines, col)
	}

	return lines
}

func trimRightSpaces(lines []string, col int) []string {
	res := make([]string, len(lines))
	for i, line := range lines {
		if col <= 0 {
			res[i] = strings.TrimRight(line, " ")
			continue
		}

		fields := strings.Split(line, "\t")
		if col > len(fields) {
			res[i] = line
			continue
		}

		fields[col-1] = strings.TrimRight(fields[col-1], " ")
		res[i] = strings.Join(fields, "\t")
	}
	return res
}

func postProcessLines(lines []string, unique bool) []string {
	if unique {
		return removeDuplicates(lines)
	}

	return lines
}

func removeDuplicates(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	result := []string{lines[0]}

	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			result = append(result, lines[i])
		}
	}

	return result
}
