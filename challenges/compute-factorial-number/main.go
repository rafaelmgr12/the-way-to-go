package main

func Factorial(n uint64) (fac uint64) {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)

}
