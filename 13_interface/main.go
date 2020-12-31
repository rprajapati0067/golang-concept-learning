package main

import (
	"fmt"
	"math"
)

type MyString string

//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}
type Cube struct {
	side float64
}

type Rectangle struct {
	width   float64
	breadth float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.breadth
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.breadth)
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func explain(i interface{}) {
	fmt.Printf("Value given to explain func is of type '%T' with value %v\n", i, i)
}

func main() {
	//var s Shape
	//s = Rectangle{5.0, 4.0}
	//r := Rectangle{5.0, 4.0}
	//
	//fmt.Println("Value of s is", s)
	//fmt.Printf("Type of s is %T\n", s)
	//fmt.Println("Area of rectangle s", s.Area())
	//fmt.Println("s == r  is", s == r)

	//ms := MyString("Hello World")
	//r := Rectangle{5., 4.5}
	//explain(ms)
	//explain(r)

	var s Shape = Cube{3}
	//	s := s.(Rectangle)  // if Type does not implement interface i.e here Rectangle doesn't implement Shape interface then we will get compile time error
	c := s.(Cube)
	value, ok := s.(Cube) // Safe way of Type assertion ok will return false if s doesn't have concrete Type or dynamic value of Type
	fmt.Println(value, ok)
	/* If Type implements the interface here s does't have a concrete value of Type because it is nil at the moment then
	Go will panic in runtime*/
	//	var s Shape
	//	c := s.(Cube)

	fmt.Println("Volume of s of interface type Shape is", c.Area())  // s static type is Area and dynamic type is Cube
	fmt.Println("Area of o of interface type Object is", c.Volume()) // o static type is Volume and dynamic type is Cube

}
