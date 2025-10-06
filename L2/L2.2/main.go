package main

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	test()
	anotherTest()
	// fmt.Println(test())
	// fmt.Println(anotherTest())
}
