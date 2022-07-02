package main

import "fmt"

func main() {
	var a = "initial" // infer the type
	fmt.Println(a)

	var b, c int = 1, 2 // multiple variables
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // zero value
	fmt.Println(e)

	f := "apple" // shorthand
	fmt.Println(f)
}
