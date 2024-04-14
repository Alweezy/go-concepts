package main

import "fmt"

func main () {
	names := [...]string{"Axa", "Alexa", "Betty", "Amin", "Eliud"}
	for _, name :=  range names {
		fmt.Printf("current name is: %s\n", name)
	}
}
