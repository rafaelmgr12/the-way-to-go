package main

import (
	"fmt"

	"github.com/rafaelmgr12/the-way-go/random-script/testing-concrete-example/even"
)

func main() {
	for i := 0; i <= 100; i++ {

		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}

}
