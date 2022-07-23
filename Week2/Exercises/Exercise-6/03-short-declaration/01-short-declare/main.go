package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Short Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
//  i: 314 f: 3.14 s: Hello b: true
// ---------------------------------------------------------
func main() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}
