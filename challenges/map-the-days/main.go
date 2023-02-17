package main

var Days = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
} // do initialization here

func findDay(n int) string {
	if n < 1 || n > 7 {
		return "Wrong Key"
	}
	return Days[n]
}
