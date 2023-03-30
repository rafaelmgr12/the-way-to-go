package main

import "fmt"

func badCall() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panicking %s\r\n", r)
		}

	}()
	badCall()
	fmt.Println("After bad call")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test() // calling test function
	fmt.Printf("Test completed\r\n")
}
