package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	side float64
}
type Rectangle struct {
	length  float64
	breadth float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (r Rectangle) getArea() float64 {
	return r.length * r.breadth
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := Square{side: 4}
	rec := Rectangle{length: 4, breadth: 5}
	printArea(sq)
	printArea(rec)
}
