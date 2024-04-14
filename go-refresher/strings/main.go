package main

import (
	"fmt"
	"strings"
)

//
//import (
//	"fmt"
//	"unicode/utf8"
//)
//
//func main() {
//	fmt.Println("This is a string")
//	fmt.Println(`This is also a string`)
//
//	// strings are immutable
//	// iterating over a string
//	data := "AB:E:E:E:EF:C:Q"
//	for _, char := range data {
//		if char != ':' {
//			fmt.Printf("%c\n", char)
//		}
//	}
//
//	fmt.Printf("The length of %s is: %d\n", data, len(data))
//	fmt.Printf("The length of %s is: %d\n", data, utf8.RuneCountInString(data))
//}

// Trimming a string
func main() {
	//string1 := "Quantum Mechanics"
	//trimmedString := strings.Trim(string1, "cs")
	//
	//fmt.Println("Trimmed string:", trimmedString)

	// Trim Left
	//trimmedString2 := strings.TrimLeft(string1, "Quan")
	//fmt.Println("Trimmed string from left:", trimmedString2)

	// Trim space

	//someRandomString := " Some white space in the beginning and end "
	//fmt.Println(someRandomString)
	//newString := strings.TrimSpace(someRandomString)
	//
	//fmt.Println(newString)

	// Let's split a string

	//data := "Some random string"
	//string2 := "Some, very, complicated, string"
	//
	//first := strings.Split(data, " ")
	//second := strings.Split(string2, ",")
	//fmt.Println(first)
	//fmt.Println(second)

	data := "Hello World, Welcome to Go Programming Language!"
	dataFields := strings.Fields(data)

	for _, field := range dataFields {
		fmt.Println(field)
	}

	fmt.Println(strings.Index(data, "e"))
	// last index
	fmt.Println(strings.LastIndex(data, "e"))
}
