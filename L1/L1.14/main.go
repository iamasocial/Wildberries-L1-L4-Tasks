package main

import "fmt"

func main() {
	arr := []interface{}{
		42,
		"wildberris",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		3.14,
	}

	for _, v := range arr {
		detectType(v)
	}
}

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Type: int, value: %v\n", v)
	case string:
		fmt.Printf("Type: string, value: %v\n", v)
	case bool:
		fmt.Printf("Type: bool, value: %v\n", v)
	case chan int:
		fmt.Println("Type: chan int")
	case chan string:
		fmt.Println("Type: chan string")
	case chan bool:
		fmt.Println("Type: chan bool")
	default:
		fmt.Printf("Unknow type: %T\n", v)
	}
}
