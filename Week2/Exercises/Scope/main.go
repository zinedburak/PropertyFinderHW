package main

import "fmt"

func main() {
	// The first x that holds the string
	x := "hello"
	for i := 0; i < len(x); i++ {
		// The second x that holds the ascii value of the x[i]
		x := x[i]
		if x != '!' {
			// Third x that has tha asci value of capital x[i]
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}
