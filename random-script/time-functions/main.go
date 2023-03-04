package main

import (
	"fmt"
	"time"
)

func Calculation() {
	var evenCount int = 0
	var oddCount int = 0
	for i := 0; i < 1000000000; i++ {
		if i%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}

	}
}
func main() {
	start := time.Now()
	Calculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Calculation took this amount of time: %s\n", delta)
}
