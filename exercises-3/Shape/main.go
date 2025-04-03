package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, width float64
}

type triangle struct {
	A, B, C float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64{
	return 2*math.Pi*c.radius
}

func(r rectangle) Area() float64{
	return r.length * r.width
}


func(r rectangle) Perimeter() float64{
	return 2*( r.length + r.width)
}

func(t triangle) Area() float64{
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func(t triangle) Perimeter() float64{
	return t.A + t.B + t.C
}

func main() {
	var s Shape

	s = circle{radius: 5}
    fmt.Println("c Area:", s.Area())
    fmt.Println("c Perimeter:", s.Perimeter())

    s = rectangle{length: 4, width: 3}
    fmt.Println("r Area:", s.Area())
    fmt.Println("r Perimeter:", s.Perimeter())

	s = triangle{A: 3, B: 4, C: 5}
	fmt.Println("t Area:", s.Area())
	fmt.Println("t Perimeter:", s.Perimeter())
}