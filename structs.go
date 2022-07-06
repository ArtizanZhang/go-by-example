package main

import "fmt"

// Go's struct are typed collection of fields.

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{
		name: name,
	}
	p.age = 42
	return &p // you can safely return a pointer to local variable as a local variable will survive the scope of the function
}

func main() {
	// new a new struct
	fmt.Println(person{"Bob", 22})

	// name the fields when initialize a  struct
	fmt.Println(person{name: "Alice", age: 21})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	//fmt.Println(person{"Fred"})   // this syntax is invalid

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// it's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	s := person{
		name: "Sean",
		age:  30,
	}

	// access struct field with a dot
	fmt.Println(s.name)

	// you can also use dots with struct pointers - the pointers are automatically de-referenced.
	sp := &s
	fmt.Println(sp.age)

	// struct are mutable.
	sp.age = 51
	fmt.Println(sp.age)
}
