package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	AnB := intersect(A, B)

	for _, v := range AnB {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func intersect(a, b []int) []int {
	set := make(map[int]struct{})

	for _, num := range a {
		set[num] = struct{}{}
	}

	result := []int{}

	for _, num := range b {
		if _, found := set[num]; found {
			result = append(result, num)
		}
	}

	return result
}
