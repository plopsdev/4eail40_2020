package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() string
	Stringer() string
}

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() string {
	return fmt.Sprintf("Area of %g", r.Width*r.Length)
}
func (c Circle) Area() string {
	return fmt.Sprintf("Area of %g", c.Radius*c.Radius*math.Pi)
}

func (r Rectangle) Stringer() string {
	return fmt.Sprintf("Rectangle of width %g and length %g", r.Width, r.Length)
}

func (c Circle) Stringer() string {
	return fmt.Sprintf("Circle of radius %g", c.Radius)
}

func Measure(s Shape) {
	fmt.Println(s.Stringer())
	fmt.Println(s.Area())
}

func main() {

	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		Measure(s)
	}
}
