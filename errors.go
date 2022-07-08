package main

import (
	"errors"
	"fmt"
)

// In Go it’s idiomatic to communicate errors via an explicit, separate return value.
// This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.
// Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can' work with 42") // errors.New constructs a basic error value with the given error message.
	}
	return arg + 3, nil // A nil value in the error position indicates that there was no error.
}

type argError struct {
	arg  int
	prob string
}

// it’s possible to use custom types as errors by implementing the Error() method on them.
// Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
func (a argError) Error() string {
	return fmt.Sprintf("arg: %d - prob: %s", a.arg, a.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{ // In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
			arg:  arg,
			prob: "can't work with 42",
		}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}

		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg, ae.prob)
	} else {
		fmt.Println("assertion:  not a argError")
	}
}
