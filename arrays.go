package main

import "fmt"

/*
author: Onur Sak
*/

func main() {
	// First way of defining an array
	// This way is kinda verbose and there is more better way (see second way)
	var array1 [3]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3

	fmt.Println(array1)
	fmt.Println(array1[0])

	// Second way of defininf an array
	// This way is more preferrable, because it is more elegant
	array2 := [3]int{1,2,3}
	
	fmt.Println(array2)
	fmt.Println(array2[0])

	// So we obtain the same functionality with few lines of code

}