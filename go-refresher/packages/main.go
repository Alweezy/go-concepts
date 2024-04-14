package main

import (
	"fmt"
	"time"
)

//func main() {
//	data := []string{"AD03A29", "Rainfall values"}
//
//	serialNumber := fmt.Sprintf("%s", data[0])
//
//	fmt.Printf("Serial number: %s\n", serialNumber)
//
//}

//func main() {
//	var name string
//	var age int
//	var isMember bool
//
//	fmt.Scan(&name)
//	fmt.Scan(&age)
//	fmt.Scan(&isMember)
//
//	fmt.Printf("Name: %s Age: %d, isMember %t", name, age, isMember)
//}

// Scanf
//func main() {
//	var name string
//	var age int
//	var isMember bool
//
//	fmt.Printf("Your Name: ")
//	fmt.Scanf("%s", &name)
//	fmt.Println("")
//	fmt.Printf("Your Age: ")
//	fmt.Scanf("%d", &age)
//	fmt.Println("")
//	fmt.Printf("Your Are you are member?: ")
//	fmt.Scanf("%t", &isMember)
//
//	fmt.Printf("Name: %s Age: %d, isMember: %t", name, age, isMember)
//}

// Scanln
//func main() {
//	var name string
//	var age int
//	var isMember bool
//
//	fmt.Scanln(&name)
//	fmt.Scanln(&age)
//	fmt.Scanln(&isMember)
//
//	fmt.Printf("Name: %s Age: %d, isMember: %t", name, age, isMember)
//}

// Sscan

//func main() {
//	var name string
//	var age int
//	var isMember bool
//
//	data, err := fmt.Sscan("Barack 12 true", &name, &age, &isMember)
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("%d elements: %s, %d, %t", data, name, age, isMember)
//}

// Errorf
func main() {
	err := fmt.Errorf("an error occurred at %v", time.Now())

	fmt.Println("An error occurred: ", err)
}
