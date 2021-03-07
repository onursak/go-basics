package main

import "fmt"

/*
author: Onur Sak
*/

func main() {

	// If the variable is initialized and not used, we get an compiler error
	// Variable declaration style 1: Static Typing
	var movie string
	movie = "Lotr"
	fmt.Println(movie)

	var durationInHours int = 3
	fmt.Println(durationInHours)

	//Variable declaration style 2: Dynamic Typing
	actor := "Elijah Wood"
	fmt.Println(actor)

	actorAge := 20
	fmt.Println(actorAge)

	// Multiple assignments at the same time
	complexNumber := complex(3, 4)
	fmt.Println(complexNumber)

	real, imaginary := real(complexNumber), imag(complexNumber)
	fmt.Println(real, imaginary)

}
