package main

import "fmt"

func main() {

	// as the fabonacci for 0 and 1 is 1 so, setting variables expilcitly as 1
	result := 1
	a := 1
	b := 1

	// calculating fabonacci seq for first 10 numbers
	for i := 0; i <= 10; i++ {
		if i > 1 { // no need to add preceding values for first two case
			result = a + b
		}
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)

		// modifying 2 values for next iteration
		a = b      // a should have the fibonacci for i-1th number for next iteration
		b = result // b should have the fibonacci for ith number for next iteration
	}
}
