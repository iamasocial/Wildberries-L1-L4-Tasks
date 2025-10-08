package main

import (
	"fmt"
	"os"
)

type F interface {
	doSome()
}

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Println()

	var err1 error
	fmt.Println(err1 == nil)
}
