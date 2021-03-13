package main

import "fmt"

/*
author: Onur Sak
*/

// Constant block
const (
	LIGHT_SPEED     = 3
	AVOGADRO_NUMBER = 20
)

// Constant block using iota
const (
	first = iota
	second = iota
)

// Constant block using iota and arithmetic
const (
	first_arithmetic = iota + 6
	second_arithmetic = iota
)

// Constant block with single statement
// iota will be applied to all
const (
	first_single = iota
	second_single
	third_single
)

// Note: iota value will be reset for every constant block

func main() {
	// Constant variables should be declared and initialized at the same time
	const PI = 3.14
	fmt.Println(PI)

	// These lines are gonna work because we didn't specify a type for constant value, so interpreter will interpret the variable dynamically every time
	fmt.Println(PI + 4)
	fmt.Println(PI + 3.2)

	// Let's define a type for the constant
	const GRAVITY int = 10
	//fmt.Println(GRAVITY + 1.2)          // We get a truncation error
	fmt.Println(float32(GRAVITY) + 1.2) // Now we can success by using conversion operation

	fmt.Println(LIGHT_SPEED, AVOGADRO_NUMBER) // It will print out 3 and 20

	fmt.Println(first, second) // It will print out 0 and 1

	fmt.Println(first_arithmetic, second_arithmetic) // It will print out 6 and 1

	fmt.Println(first_single, second_single, third_single) // It will print out 0 and 1 and 2

}
