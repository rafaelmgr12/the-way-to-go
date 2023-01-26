package main

func fibarray(term int) []int {

	farr := make([]int, term)

	if term == 1 {
		farr[0] = 0
		return farr
	}
	farr[0], farr[1] = 0, 1

	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}

	return farr
}
