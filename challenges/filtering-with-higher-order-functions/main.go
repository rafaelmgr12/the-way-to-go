package main

func Filter(s []int, fn func(int) bool) []int {

	var r []int
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r

}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
