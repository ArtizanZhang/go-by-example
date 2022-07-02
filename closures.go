package main

import "fmt"

// returns another function, which we define anonymously in the body of intSeq,
// the returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println("NEW INTS")
	fmt.Println(newInts())
}
