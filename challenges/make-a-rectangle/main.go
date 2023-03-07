package main

type Rectangle struct { // struct of type Rectangle
	length, width int
}

func (r *Rectangle) Area() int { // method calculating area of rectangle

	return r.length * r.width
}

func (r *Rectangle) Perimeter() int { // method calculating perimeter of rectangle

	return 2*r.length + 2*r.width
}
