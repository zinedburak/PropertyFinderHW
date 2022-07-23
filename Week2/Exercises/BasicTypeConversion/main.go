package main

import "fmt"

func main() {
	var x = 46
	var y = 34.0
	var z int8

	// If you want to see the compile error uncomment the line below
	//fmt.Println(x + y + z)

	fmt.Println(float64(x) + y + float64(z))
	fmt.Println(int(y) + x + int(z))
}
