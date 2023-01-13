package main

type flt func(int) bool // aliasing type

func isEven(n int) bool { // check if n is even or not
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(sl []int, f flt) (yes, no []int) { // split s into two slices: even and odd
	for _, val := range sl {
		if f(val) {
			yes = append(yes, val)
		} else {
			no = append(no, val)
		}
	}
	return
}
