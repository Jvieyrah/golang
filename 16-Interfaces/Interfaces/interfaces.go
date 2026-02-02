package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(f form) float64 {
	return f.area()
}

func main() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}

	fmt.Println(getArea(r))
	fmt.Println(getArea(c))
}
