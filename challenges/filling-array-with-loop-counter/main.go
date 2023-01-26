package main

import "fmt"

func main() {
	var arr [15]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)
}
