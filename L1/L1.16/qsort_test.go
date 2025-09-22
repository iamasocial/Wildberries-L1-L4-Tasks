package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuckSort(t *testing.T) {
	tests := []struct {
		input, expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{8, 8, 8, 8}, []int{8, 8, 8, 8}},
		{[]int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1}},
		{[]int{-9, 5, 2, 0, -3, -6, 11}, []int{-9, -6, -3, 0, 2, 5, 11}},
	}

	for _, test := range tests {
		got := quickSort(test.input)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("quickSort(%v) got: %v, expected: %v\n", test.input, got, test.expected)
		}
	}
}

func BenchmarkQuckSort(b *testing.B) {
	arr := make([]int, 100000)
	for i := range arr {
		arr[i] = rand.Intn(1000000)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		quickSort(arr)
	}
}
