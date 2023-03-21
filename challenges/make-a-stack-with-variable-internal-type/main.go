package main

import (
	"fmt"

	"github.com/rafaelmgr12/the-way-go/challenges/make-a-stack-with-variable-internal-type/mystack"
)

var st1 mystack.Stack

func main() {
	st1.Push("Brown")
	st1.Push(3.14)
	st1.Push(100)
	st1.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
