package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type  Rectangle struct{
	Width float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return  2 * (r.Width + r.Length)
}
func  PrintShapeArea(s Shape )  {
	fmt.Println("Area is : ", s.Area())
}

func PrintShapePerimeter(s Shape)  {
	fmt.Println("Perimeter is ", s.Perimeter())
}

func main(){
	r := Rectangle{Width: 5, Length: 20}
	PrintShapeArea(r)
	PrintShapePerimeter(r)
}
//go run mytest.go
