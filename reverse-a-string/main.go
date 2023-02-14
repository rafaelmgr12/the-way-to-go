package main

func reverse(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r

}
