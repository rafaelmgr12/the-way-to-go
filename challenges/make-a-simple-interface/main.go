package main

import "fmt"

type Simpler interface { // interface implementing functions called on struct Simple
	Get() int
	Set(u int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {

	return p.i
}

func (p *Simple) Set(u int) {

	p.i = u
}

func fI(it Simpler) int { // function calling both methods through interface
	it.Set(5)
	return it.Get()
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with a receiver type pointer
}
