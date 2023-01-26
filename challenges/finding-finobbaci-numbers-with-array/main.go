package main

var fib [10]int64

func fibs() [10]int64 {
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < 10; i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}
	return fib
}
