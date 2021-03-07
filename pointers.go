package main

import "fmt"

/*
author: Onur Sak
*/

func main() {
	var name *string
	fmt.Println(name) // It prints nil because currently not pointing to anything

	// Dereferencing
	// Grab that data through the pointer
	// *name = "Thranduil" // It gives an invalid memory or nil pointer dereference error, because we are currently not pointing to a valid memory address

	name = new(string) // Allocate a memory for string data type
	fmt.Println(name)  // It prints the memory address of the "name" that starts like 0x...

	// Since we allocated a memory, then we can store the data where the pointer points out
	*name = "Thranduil"
	fmt.Println(*name) // Now it can print the Thranduil sucessfully

	// Note: Go doesn't support pointer arithmetic to keep programs as reliable and predictable as possible. It is consistent with the Go's way of thinking: Simplicity

	bestTvSeries := "Lost" // Is there any doubt for this? :D Anyway, let's get back to the work!
	// Now I want to get the memory address of the bestTvSeries variable, we can use the address of operator

	pointerBestTvSeries := &bestTvSeries
	fmt.Println(pointerBestTvSeries, *pointerBestTvSeries) // It prints the memory location of the bestTvSeries, and the value of the bestTvSeries variable

	// Now assign a new value to bestTvSeries
	bestTvSeries = "Sherlock"
	// Let's print the memory address and dereference again
	fmt.Println(pointerBestTvSeries, *pointerBestTvSeries) // It still prints the same memory address, we changed the variable's value however it is still stored in the same memory address

}
