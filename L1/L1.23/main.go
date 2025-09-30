package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}
	i := 2

	fmt.Println("Original int slice:", ints)

	////////////////////////////////
	ints, ok := RemoveIntAt(ints, i)
	fmt.Println(ints, ok)

	ints, ok = RemoveIntAt(ints, 10)
	fmt.Println(ints, ok)
	////////////////////////////////

	ints = []int{1, 2, 3, 4, 5}

	////////////////////////////////
	ints, ok = RemoveIntFast(ints, i)
	fmt.Println(ints, ok)

	ints, ok = RemoveIntFast(ints, 10)
	fmt.Println(ints, ok)
	////////////////////////////////

	a, b, c, d, e := "a", "b", "c", "d", "e"
	objs := []*string{&a, &b, &c, &d, &e}
	i = 1

	fmt.Println("Original pointers slice", objs)

	////////////////////////////////
	objs, ok = RemovePtrAt(objs, i)
	fmt.Println(objs, ok)

	objs, ok = RemovePtrAt(objs, 10)
	fmt.Println(objs, ok)
	////////////////////////////////

	objs = []*string{&a, &b, &c, &d, &e}

	////////////////////////////////
	objs, ok = RemovePtrFast(objs, i)
	fmt.Println(objs, ok)

	objs, ok = RemovePtrFast(objs, 10)
	fmt.Println(objs, ok)
}

func RemoveIntAt(s []int, i int) ([]int, bool) {
	if i < 0 || i > len(s)-1 {
		return s, false
	}

	return append(s[:i], s[i+1:]...), true
}

func RemovePtrAt[T any](s []*T, i int) ([]*T, bool) {
	if i < 0 || i > len(s)-1 {
		return s, false
	}

	s[i] = nil
	return append(s[:i], s[i+1:]...), true
}

func RemoveIntFast(s []int, i int) ([]int, bool) {
	if i < 0 || i > len(s)-1 {
		return s, false
	}

	s[i] = s[len(s)-1]
	return s[:len(s)-1], true
}

func RemovePtrFast[T any](s []*T, i int) ([]*T, bool) {
	if i < 0 || i > len(s)-1 {
		return s, false
	}

	s[i] = s[len(s)-1]
	s[len(s)-1] = nil
	return s[:len(s)-1], true
}
