package main

func sumInts(list ...int) (sum int) {
	for _, v := range list {
		sum += v
	}
	return
}
