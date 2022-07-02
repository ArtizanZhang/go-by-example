package main

import (
	"fmt"
	"math"
)

const s string = "constant" // constant statement can appear anywhere a var statement can
func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 2e20 / n
	fmt.Println(d) // perform arithmetic with arbitrary precision

	// a numeric constant has no type until it's required in a context, such as by an explicit conversion or variable assignment or function call
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
