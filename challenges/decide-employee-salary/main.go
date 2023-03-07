package main

/* basic data structure upon which we'll define methods */
type employee struct {
	salary float32
}

/*
a method which will add a specified percent to an

	employees salary
*/
func (this *employee) giveRaise(pct float32) {
	this.salary = this.salary * (1 + pct)
	return
}
