package main

import "fmt"

// Go supports embedding of structs and interfaces to express a more seamless composition of types
// this is not to be confused with // go:embed which is a go directive introduced in Go Version 1.16+
// to embed files and folders into the application binary.

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base // A	container embeds a base.  an embedding likes a field without a name.
	str  string
}

func main() {

	// when creating structs with literals, we have to initialize the embedding explicitly;
	co := container{
		base: base{ //  here the embedded type  serves as the field name
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}", co.num, co.str) // We can access the base’s fields directly on co, e.g. co.num.
	fmt.Println("also num: ", co.base.num)              // Alternatively, we can spell out the full path using the embedded type name.

	// since container embeds base, the methods of base also become methods of a container
	fmt.Println("describe: ", co.describe())                   // here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe path to base: ", co.base.describe()) // alternatively, we can spell out the full path for methods in base too.

	type describer interface {
		describe() string
	}

	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	// Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describer: ", d.describe())
}
