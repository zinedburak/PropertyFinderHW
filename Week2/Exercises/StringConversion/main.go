package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s := "My name is Burak Deniz and I' am "
	age := 25

	// If you want to see the compile error uncomment the line below
	//fmt.Println(s + age)
	ageText := strconv.Itoa(age)
	fmt.Println(s + ageText)
	fmt.Println(reflect.TypeOf(ageText))

	ageInt, _ := strconv.ParseInt(ageText, 10, 64)
	fmt.Println(reflect.TypeOf(ageInt))

}
