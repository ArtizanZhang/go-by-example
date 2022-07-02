package main

import "fmt"

func main() {

	// an array hold exactly 5 ints. the type of elements and length are both part of the array's type
	var a [5]int
	fmt.Println("emp:", a) // zero-valued, which for ints means 0s.

	a[4] = 100 // set a value
	fmt.Println("set:", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len", len(a)) // builtin len

	b := [5]int{1, 2, 3, 4, 5} // declare and initialize an array in one line
	fmt.Println("dcl:", b)

	// compose types to build multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
