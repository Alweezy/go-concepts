package main

import (
	"fmt"
	"time"
)

// application has more than 1 thing at the same time
// Goroutines become inevitable in concurrency and parallelism
// These are Golang wrapper on top of threads created and managed by
// Go runtime rather than the underlying os

// Go Lang achieves concurrency using Goroutines:
// They allow us to run multiple methods or functions
// concurrently in the same address space inexpensively

// All goroutines are managed by the main goroutine
// If the main goroutine terminated for some reason,
// all the other goroutines would terminate

//goroutines run in the backend

// https://youtu.be/zh7d-9O_Hd0?t=10304

// creating go routines adding keyword go to a statement, function

// main function in the package main is the main goroutine

//func main() {
//	go display()
//	fmt.Println("Hello")
//
//	// timeout to allow goroutine
//	time.Sleep(1 * time.Second)
//	fmt.Println("GoodBye!")
//}
//
//func display() {
//	fmt.Println("Go routine ...")
//}

// goroutines execute independently
//func main() {
//	go hello()
//	fmt.Println("Hello from main func ...")
//	time.Sleep(1 * time.Second)
//	fmt.Println("GoodBye 😉")
//}
//
//func hello() {
//	go hello2()
//	fmt.Println("Goroutine 1 ...")
//}
//
//func hello2() {
//	fmt.Println("Goroutine 2 ...")
//}

//func main() {
//	fmt.Println("Main goroutine starts here ...")
//	go name()
//
//	go id()
//	time.Sleep(5 * time.Second)
//	fmt.Println("\nMain goroutine ends here ...")
//}
//
//func name() {
//	array1 := [4]string{"Jesica", "Elba", "Deen", "Ali"}
//	for _, item := range array1 {
//		time.Sleep(150 * time.Millisecond)
//		fmt.Print("...")
//		time.Sleep(150 * time.Millisecond)
//		fmt.Print("...")
//		time.Sleep(150 * time.Millisecond)
//		fmt.Print("...")
//		time.Sleep(150 * time.Millisecond)
//		fmt.Printf("%s\n", item)
//	}
//}
//
//func id() {
//	array2 := [4]int{301, 304, 405, 200}
//
//	for _, _id := range array2 {
//		time.Sleep(3 * time.Second)
//		fmt.Printf("%d\n", _id)
//	}
//}

// anonymous goroutine
func main() {
	fmt.Println("Hello to main function")

	go func() {
		fmt.Println("Welcome to goroutine")
	}()

	time.Sleep(1 * time.Second)
	fmt.Print("Goodbye! ")
}
