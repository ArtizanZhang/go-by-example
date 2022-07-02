package main

import "fmt"

// Go requires explicit returns. i.e. it won't automatically return the value of the last expression.
func plus(a int, b int) int {
	return a + b
}

// when you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter the declares the type .
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// call a function just as you'd expect
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
