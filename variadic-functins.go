package main

import "fmt"

// variadic functions can be called with any number of trailing arguments.
// for example. fmt.Println is a common variadic function.

func sum(nums ...int) {
	fmt.Print(nums, " ") // within the function, the type of nums is equivalent to []int, we can call len(nums), iterate over it with range, etc.

	fmt.Println("len: ", len(nums))
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4} //  if you already have multiple args in a slice,  apply them to a variadic function using func(slice...) like this,
	sum(nums...)
}
