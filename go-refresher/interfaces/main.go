package main

import "fmt"

//type Rectangle interface {
//	Perimeter() float64
//	Area() float64
//}
//
//type Compute struct {
//	length float64
//	width  float64
//}
//
//func (c Compute) Perimeter() float64 {
//	return 2 * (c.length + c.width)
//}
//
//func (c Compute) Area() float64 {
//	return c.length * c.width
//}
//
//func main() {
//
//	//rectangle := Rectangle(
//	//	Compute{20, 10},
//	//)
//
//	var rectangle Rectangle
//
//	rectangle = Compute{100, 20}
//
//	area := rectangle.Area()
//	perimeter := rectangle.Perimeter()
//
//	fmt.Println("Area: ", area)
//	fmt.Println("Perimeter: ", perimeter)
//}

// Empty interfaces

//func main() {
//	name := "Hello"
//	number := 10.64
//	typeOf(name)
//	typeOf(number)
//}
//
//func typeOf(arg interface{}) {
//	fmt.Printf("The type of %v is: %T\n", arg, arg)
//}

// Type assertions for interfaces

//func main() {
//	name := "Bob"
//	typeOf(name)
//
//	number := 123.4
//	typeOf(number)
//}
//
//func typeOf(arg interface{}) {
//	value, ok := arg.(string)
//	fmt.Println("Status ..... ", ok)
//	if !ok {
//		fmt.Printf("Expected %v to be string, found %T\n", arg, arg)
//	} else {
//		fmt.Printf("Value of %v is as expected: %T\n", value, arg)
//	}
//}

// Type switch

func main() {
	name := "Bee"
	number := 123
	decimal := 125.45

	typeOf(name)
	typeOf(number)
	typeOf(decimal)
}

func typeOf(arg interface{}) {
	switch arg.(type) {
	case int:
		fmt.Printf("Type Int: %v is type %T\n", arg, arg)

	case float64, float32:
		fmt.Printf("Type Float: %v is type %T\n", arg, arg)

	case string:
		fmt.Printf("Type String: %v is type %T\n", arg, arg)
	default:
		fmt.Printf("Unexpected data type ... %T\n", arg)
	}
}
