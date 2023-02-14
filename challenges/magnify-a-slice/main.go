package main

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	for i, v := range s {
		for j := 0; j < factor; j++ {
			ns[i*factor+j] = v
		}
	}
	return ns
}

// func enlarge(s []int, factor int) []int {
// 	ns := make([]int, len(s) * factor)
// 	copy(ns, s)
// 	return ns
// }
