package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int64("num", 0, "number to modify")
	pos := flag.Uint("pos", 0, "position of bit to change (0-indexed, right to left)")
	value := flag.Bool("value", false, "value to set the bit (true for 1, false for 0)")
	flag.Parse()

	fmt.Printf("Original number: %d (binary: %b)\n", *num, *num)
	newNum := setBit(*num, *pos, *value)
	fmt.Printf("Modified number: %d (binary: %b)\n", newNum, newNum)

}

func setBit(num int64, pos uint, value bool) int64 {
	if value {
		return num | (1 << pos)
	}

	return num &^ (1 << pos)
}
