package v2

import "math"

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

// methodlar structlarla birlikte kullanılır. Kullanim yapisi fonksiyonlara cok benze
// func (receiverName ReceiverType) MethodName(args)
// receiverName: methodun bagli oldugu structun ismi ve structerin bas harfi kucuk olmasi bir kuraldir

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
