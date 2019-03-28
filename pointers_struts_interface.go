package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *Circle) area() float64 {
	return math.Pi * (*c).radius * (*c).radius

}

func info(s shape) {
	fmt.Println("the area of the shape is ", s.area())
}

func main() {

	c := Circle{
		radius: 24.0,
	}
	fmt.Printf("%T\n", &c)
	info(&c)
	fmt.Println("The final value of c is", c)
	temp := &c
	temp.radius = 123
	fmt.Println("The new value of c is", c)
	fmt.Printf("%T\n", &c)
	info(&c)
	fmt.Println("The  new final value of c is", c)
}
