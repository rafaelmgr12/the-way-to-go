package even

func Even(i int) bool { // Exported functions
	return i%2 == 0
}

func Odd(i int) bool {
	return i%2 != 0
}
