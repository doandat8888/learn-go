package main

import "fmt"

type shape interface {
	calcArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{}
	sq.sideLength = 5
	printArea(sq)
	tr := triangle{}
	tr.base = 4
	tr.height = 10
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.calcArea())
}

func (sq square) calcArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) calcArea() float64 {
	return 0.5 * tr.base * tr.height
}
