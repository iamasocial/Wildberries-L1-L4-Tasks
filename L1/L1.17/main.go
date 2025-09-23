package main

import "fmt"

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 8
	fmt.Printf("index of %d is %d\n", target, binarySearch(sortedArr, target))
}

func binarySearch(sortedArr []int, target int) int {
	begin, end := 0, len(sortedArr)-1

	for begin <= end {
		mid := begin + (end-begin)/2

		if target == sortedArr[mid] {
			return mid
		}

		if target > sortedArr[mid] {
			begin = mid + 1
			continue
		}

		if target < sortedArr[mid] {
			end = mid - 1
			continue
		}
	}

	return -1
}
