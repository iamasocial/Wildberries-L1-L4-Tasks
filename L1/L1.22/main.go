package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	MaxInt = int(^uint(0) >> 1) //  9223372036854775807 int64 //  2147483647 int32
	MinInt = -MaxInt - 1        // -9223372036854775808 int64 // -2147483648 int32
)

func main() {
	var aStr, bStr string
	fmt.Printf("Enter first number\n> ")
	fmt.Scan(&aStr)
	fmt.Printf("Enter second number\n> ")
	fmt.Scan(&bStr)

	aInt, aErr := strconv.Atoi(aStr)
	bInt, bErr := strconv.Atoi(bStr)

	useBig := aErr != nil || bErr != nil

	aBig, aOk := new(big.Int).SetString(aStr, 10)
	bBig, bOk := new(big.Int).SetString(bStr, 10)

	if !aOk || !bOk {
		fmt.Printf("Error: incorrect numbers entered\n")
		return
	}

	if useBig || addOverflow(aInt, bInt) {
		add := new(big.Int).Add(aBig, bBig)
		fmt.Printf("big.Int:\t%s + %s = %s\n", aBig, bBig, add)
	} else {
		add := aInt + bInt
		fmt.Printf("int:\t%d + %d = %d\n", aInt, bInt, add)
	}

	if useBig || subOverflow(aInt, bInt) {
		sub := new(big.Int).Sub(aBig, bBig)
		fmt.Printf("big.Int:\t%s - %s = %s\n", aBig, bBig, sub)
	} else {
		sub := aInt - bInt
		fmt.Printf("int:\t%d - %d = %d\n", aInt, bInt, sub)
	}

	if useBig || mulOverflow(aInt, bInt) {
		mul := new(big.Int).Mul(aBig, bBig)
		fmt.Printf("big.Int:\t%s * %s = %s\n", aBig, bBig, mul)
	} else {
		mul := aInt * bInt
		fmt.Printf("int:\t%d * %d = %d\n", aInt, bInt, mul)
	}

	if bInt == 0 {
		fmt.Printf("Error: division by zero\n")
		return
	}

	if useBig || divOverflow(aInt, bInt) {
		div := new(big.Int).Div(aBig, bBig)
		fmt.Printf("big.Int:\t%s / %s = %s\n", aBig, bBig, div)
	} else {
		div := aInt / bInt
		fmt.Printf("int:\t%d / %d = %d\n", aInt, bInt, div)
	}
}

func addOverflow(a, b int) bool {
	return (b > 0 && a > MaxInt-b) || (b < 0 && a < MinInt-b)
}

func subOverflow(a, b int) bool {
	return (b > 0 && a < MinInt+b) || (b < 0 && a > MaxInt+b)
}

func mulOverflow(a, b int) bool {
	return (a > 0 && b > 0 && a > MaxInt/b) ||
		(a < 0 && b < 0 && a < MaxInt/b) ||
		(a > 0 && b < 0 && a > MinInt/b) ||
		(a < 0 && b > 0 && a < MinInt/b)
}

func divOverflow(a, b int) bool {
	return (a == MinInt) && (b == -1)
}
