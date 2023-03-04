package main

import "fmt"

// make struct here
type challenge struct {
	number float64
	int
	string
}

func main() {
	// create struct using struct literal and print its field using fmt package
	c := challenge{number: 1.0, int: 2, string: "three"}
	fmt.Println(c.number, c.int, c.string)
}
