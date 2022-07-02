package main

import "fmt"

// range iterates over elements in a variety of data structures
func main() {
	nums := []int{2, 3, 4}
	// range to sum the numbers in a slice, arrays work like this too.
	sum := 0
	for _, num := range nums { // ignored it with the blank identifier _.
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range on map iterate over k-v pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range on strings iterate over Unicode code points. the first value is the starting byte index of the rune
	// and the second the rune itself.
	// https://gobyexample.com/range#:~:text=rune%20itself.%20See-,Strings%20and%20Runes,-for%20more%20details
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

/**
output:
		sum: 9
		index: 1
		a -> apple
		b -> banana
		key: a
		key: b
		0 103
		1 111
*/