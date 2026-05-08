package main

import (
	"fmt"
)

type area interface {
	AreaCalculator(l, b int) int
}

type squareDimention struct {
	length int
	width  int
}

type rectangleDimention struct {
	length int
	width  int
}

type triangleDimention struct {
	length int
	width  int
}

func main() {
	sq := squareDimention{5, 5}
	print(sq, 0, 0, "Square")
	re := rectangleDimention{}
	print(re, 9, 3, "Rectangle")
	tr := triangleDimention{}
	print(tr, 20, 3, "Triangle")
}

func (sl squareDimention) AreaCalculator(l int, w int) int {
	return sl.length * sl.width
}

func (re rectangleDimention) AreaCalculator(l int, w int) int {
	return l * w
}

func (tr triangleDimention) AreaCalculator(l int, b int) int {
	return (l * b) / 2
}

func print(a area, l int, b int, shape string) {
	out := a.AreaCalculator(l, b)
	fmt.Printf("\nThe Area of %s is %v", shape, out)
}
