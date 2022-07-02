package main

import "fmt"

func main() {
	// unlike arrays, slices are typed only by the elements they contain, not the length of elements
	s := make([]string, 3, 3)
	fmt.Println("emp: ", s)

	// set like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println("len: ", len(s))

	// slice support several more that make them richer than arrays
	// append , note that we need to accept a return value from append as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	s = append(s, []string{"g", "h"}...)
	fmt.Println("apd:", s)

	// copy ,  create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// support  slice operator
	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl2；", l)

	// declare and initialize for slice in a single line
	t := []string{"aa", "bb"}
	fmt.Println("dcl:", t)

	// multi-dimensional data structures
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
