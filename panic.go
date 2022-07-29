package main

import "os"

// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

func main() {
	panic("a problem")

	/**
	  When first panic in main fires, the program exits without reaching the rest of the code.
	  If you’d like to see the program try to create a temp file, comment the first panic out.
	*/
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/*
Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible
*/
