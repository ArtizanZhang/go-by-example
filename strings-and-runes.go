package main

import (
	"fmt"
	"unicode/utf8"
)

//  Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
//  In other languages, strings are made of “characters”.
//  In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
//  This Go blog post is a good introduction to the topic.
func main() {

	const s = "สวัสดี"

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("len: ", len(s)) // 18

	// Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s)) // 6

	for idx, runeValue := range s {
		fmt.Printf("%#U start at %d\n", runeValue, idx) // %#U
		//fmt.Println(runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s)
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue) // passing  a rune value to a function
	}

}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
