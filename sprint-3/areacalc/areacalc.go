package areacalc

import "bytes"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	a    float64
	b    float64
	name string
}

func NewRectangle(a_ float64, b_ float64, name_ string) *Rectangle {
	return &Rectangle{a: a_, b: b_, name: name_}
}

func (rec Rectangle) Area() float64 {
	return rec.a * rec.b
}

func (rec Rectangle) Type() string {
	return rec.name
}

type Circle struct {
	radius float64
	name   string
}

func (cir Circle) Area() float64 {
	return pi * cir.radius * cir.radius
}

func (cir Circle) Type() string {
	return cir.name
}

func NewCircle(radius_ float64, name_ string) *Circle {
	return &Circle{radius: radius_, name: name_}
}

func AreaCalculator(figures []Shape) (string, float64) {
	var area float64
	var buf bytes.Buffer

	for i := range figures {
		if i != 0 {
			buf.WriteString("-")
		}
		buf.WriteString(figures[i].Type())
		area += figures[i].Area()
	}

	return buf.String(), area
}
