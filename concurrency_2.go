package main

import "fmt"
import "math"

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func info(s Shape) {
	fmt.Println("Area =", s.Area())
}

func main() {

	c := Circle{
		Radius: 5.0,
	}

	info(&c)
	// magically,  below works. It works for struct but not for interfaces
	fmt.Println(c.Area())

}
