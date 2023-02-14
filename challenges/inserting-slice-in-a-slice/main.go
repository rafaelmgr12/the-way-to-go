package main

// slice is the slice to which another slice will be added
// insertion is the slice that needs to be inserted.
// index is the position for insertion
func insertSlice(slice, insertion []string, index int) []string {

	return append(slice[:index], append(insertion, slice[index:]...)...)

}
