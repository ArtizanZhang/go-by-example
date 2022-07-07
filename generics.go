package main

import "fmt"

// starting with version 1.18. Go has added support for generics, also known as type parameters

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}
type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func main() {

	var m = map[int]string{1: "2", 2: "2"}
	fmt.Println("keys m:", MapKeys(m))

	// When invoking generic functions, we can often rely on type inference.
	// Note that we don’t have to specify the types for K and V when calling MapKeys - the compiler infers them automatically.
	fmt.Println("keys specify type explicitly m:", MapKeys[int, string](m))

	lst := List[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)

	fmt.Println("list: ", lst.GetAll())
}
