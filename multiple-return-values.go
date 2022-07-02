package main

import "fmt"

// Go has builtin support for multiple return values. this feature is used often in idiomatic Go,
// for example to return both result and error values form a function .

// return 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals() // if you only want a subset of the returned values, use the blank identifier _ .
	fmt.Println(c)
}
