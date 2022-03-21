package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type Rect struct {
	width  float64
	length float64
}

func (r *Rect) area() float64 {
	return r.width * r.length
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius
}

func main() {
	c := Circle{4}
	r := Rect{4, 6}

	shapes := []shape{&c, &r}

	fmt.Println(shapes[1].area())
}
