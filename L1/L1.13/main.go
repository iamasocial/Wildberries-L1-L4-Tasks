package main

import "fmt"

func main() {
	a := 5  // 0101
	b := 10 // 1010

	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)

	a, b = swapAddSub(a, b)
	fmt.Printf("After swapAddSub: a=%d, b=%d\n", a, b)

	a, b = swapXOR(a, b)
	fmt.Printf("After swapXOR: a=%d, b=%d\n", a, b)
}

func swapAddSub(a, b int) (int, int) {
	// a = a + b = 5 + 10 = 15
	// b = a - b = 15 - 10 = 5
	// a = a - b = 15 - 5 = 10
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func swapXOR(a, b int) (int, int) {
	// a = a xor b = 0101 ^ 1010 = 1111 (15)
	// b = a xor b = 1111 ^ 1010 = 0101 (5)
	// a = a xor b = 1111 ^ 0101 = 1010 (10)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func swapMulDiv(a, b int) (int, int) { // unsafe
	// a = a * b = 5 * 10 = 50
	// b = a / b = 50 / 10 = 5
	// a = a / b = 50 / 5 = 10
	if a == 0 || b == 0 {
		fmt.Printf("Division by zero error\n")
		return a, b
	}
	a = a * b
	b = a / b
	a = a / b
	return a, b
}

// func swapTemp(a, b int) (int, int) {
// 	a, b = b, a
// 	return a, b // golang creates temp variable under the hood anyway
// }
