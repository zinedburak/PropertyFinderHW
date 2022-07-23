package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius int
}

type Wheel struct {
	Circle
	hello string
}

func helloworld(s string) (returnValue, emptyString string) {
	s = "Burak Deniz"
	returnValue = s
	return
}

func main() {
	returnValue, emptyString := helloworld("Hello")
	fmt.Println(returnValue, emptyString)
}
