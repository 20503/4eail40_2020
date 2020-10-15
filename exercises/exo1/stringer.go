package main

import "fmt"

//Rectangle a
type Rectangle struct {
	width, height int
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Square of width %d and height %d", r.width, r.height)
}

//Circle a
type Circle struct {
	radius int
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", c.radius)
}

//Shape a
type Shape interface {
	Println() string
}

func (r Rectangle) Println() string {
	return r.String()
}

func (c Circle) Println() string {
	return c.String()
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
