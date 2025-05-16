package main

import "fmt"

func someFunction() {
	fmt.Println("panic before")
	panic("panic")
	fmt.Println("panic after")

}

func main() {
	defer func() {
		s := recover()
		fmt.Println("catch panic", s.(string))
	}()

	someFunction()
	fmt.Println("hello world")
}
