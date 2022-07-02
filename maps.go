package main

import "fmt"

// maps are Go's builtin associative data type. (sometimes class hashes or dicts in other languages)
func main() {
	m := make(map[string]int) // use builtin make to create an empty map

	// set k-v pairs
	m["k1"] = 1
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"] // get a value for a key
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m)) // builtin len returns the number of k-v pairs

	delete(m, "k2") // builtin delete removes k-v pairs from a map

	// indicates if the key was present in the map,
	// this can be used to disambiguate between missing keys and keys with zero values like 0 or ""
	value, prs := m["k2"]
	fmt.Println("value: ", value) // 0 (zero value)
	fmt.Println("prs: ", prs)     // false

	// declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}
