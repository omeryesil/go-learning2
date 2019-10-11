// Assignment:
// 	Create structs for triangle and square
//	Add functions to triangle and square that calculates the area (use receivers)
//	Create an interface that prints the area of triangle/square

package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		height: 4,
		base:   3,
	}

	s := square{
		sideLength: 10,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(reflect.TypeOf(s), ":", s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * 0.5 * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}