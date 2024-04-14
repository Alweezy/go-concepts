package main

func main() {
	//simpleFunction()
	//computeSum(20, 40)
	//sum := computeSum(20, 50)
	//fmt.Printf("x + y = %d", sum)

	// function as value
	//squareRoot := func(x float64) float64 {
	//	return math.Sqrt(x)
	//}
	//
	//fmt.Println(squareRoot(9))

	//area, perimeter := areaAndPerimeter(20, 30)
	//fmt.Printf("Area: %d\t Perimeter: %d", area, perimeter)

	// call by value
	//a := 100
	//b := 200
	//
	//fmt.Printf("Before swap a: %d\t b: %d\n", a, b)
	//swap(a, b)
	//fmt.Printf("After swap a: %d\t b: %d\n", a, b)
	//
	//fmt.Printf("Before swap by reference a: %d\t b: %d\n", a, b)
	//swapByReference(&a, &b)
	//fmt.Printf("After swap a by reference: %d\t b: %d\n", a, b)

	// anonymous function, functions with no names
	//func() {
	//	fmt.Println("Hello anonymous function")
	//}() // make a call to the anonymous function

	// assign anonymous function to a variable
	//result := func() {
	//	fmt.Println("Anonymous function assigned to variable")
	//}
	//
	//result()

	// anonymous function with args

	//func(name string) {
	//	fmt.Printf("Hello %s", name)
	//}("Alex")

	//shoutOut := func(name string) string {
	//	return "Hello " + name
	//}
	//
	//fmt.Printf(shoutOut("Alex"))

}

//func simpleFunction() {
//	fmt.Println("Hello World!")
//}

//func computeSum(x, y int) {
//	total := x + y
//	fmt.Printf("x + y = %d", total)
//}

//functions with return types

//func computeSum(x, y int) int {
//	return x + y
//}

// multiple return values
//func areaAndPerimeter(length int, width int) (area int, perimeter int) {
//	perimeter = 2 * (length + width)
//	area = length * width
//
//	return area, perimeter
//}

// deals with actual values
//func swap(x, y int) int {
//	temp := x
//	x = y
//	y = temp
//	return temp
//}

// call by reference, copies memory address
//func swapByReference(x, y *int) int {
//	temp := *x
//	*x = *y
//	*y = temp
//
//	return temp
//}
