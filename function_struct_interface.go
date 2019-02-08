package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (c *Circle) area() float64 {
	return math.Pi * (*c).radius * (*c).radius

}

func (sq *Square) area() float64 {
	return (*sq).length * (*sq).length
}

func info(s Shape) {
	fmt.Println(s.area())
}

func main() {

	cir1 := Circle{
		radius: 123.1,
	}

	sq1 := Square{
		length: 11.0,
	}

	info(&cir1)
	info(&sq1)

}
