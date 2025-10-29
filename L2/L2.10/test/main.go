package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "apple" < "apple ")
}

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// var humanReadableSuffixes = map[string]int{
// 	"K": 1024, "M": 1024 * 1024,
// 	"G": 1024 * 1024 * 1024, "T": 1024 * 1024 * 1024 * 1024,
// }

// func main() {
// 	input := "1540K"
// 	v, err := parseHuman(input)
// 	fmt.Printf("value: %d error: %v\n", v, err)
// }

// func parseHuman(s string) (int, error) {
// 	if len(s) == 0 {
// 		return 0, fmt.Errorf("empty string)")
// 	}

// 	mult := 1
// 	suffix := s[len(s)-1:]
// 	fmt.Println(suffix)

// 	if v, ok := humanReadableSuffixes[suffix]; ok {
// 		s = s[:len(s)-1]
// 		fmt.Println(s)
// 		mult = v
// 	}

// 	value, err := strconv.ParseFloat(s, 64)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(value * float64(mult)), nil
// }

// func compareHumanNumeric(a, b string) int {
// 	ah, err1 := parseHuman(a)
// 	bh, err2 := parseHuman(b)
// 	if err1 != nil || err2 != nil {
// 		return strings.Compare(a, b)
// 	}

// 	if ah < bh {
// 		return -1
// 	}

// 	if ah == bh {
// 		return 0
// 	}

// 	return 1
// }
