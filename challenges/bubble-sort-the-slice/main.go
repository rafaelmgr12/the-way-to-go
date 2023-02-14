package main

func bubbleSort(sl []int) []int {

	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl)-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	return sl

}
