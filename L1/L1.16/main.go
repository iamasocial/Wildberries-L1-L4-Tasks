package main

import "fmt"

func main() {
	arr := []int{5, 4, 8, 1, 4, 2, 3}
	fmt.Printf("original array: %v\n", arr)
	fmt.Printf("sorted array: %v\n", quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}

	for i, v := range arr {
		if i == len(arr)/2 {
			continue
		}

		if v < pivot {
			left = append(left, v)
			continue
		}

		if v >= pivot {
			right = append(right, v)
			continue
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}
