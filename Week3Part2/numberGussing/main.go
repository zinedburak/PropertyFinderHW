package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	game()

}
func game() {
	secretInt, decimalMap := createRandomInt()
	fmt.Println(decimalMap)
	var enteredValue string
	fmt.Scanln(&enteredValue)
	for enteredValue != secretInt {
		fmt.Println("Wrong Guess Please Enter a New Guess")
		positionAndValue := 0
		justValue := 0
		for i := 0; i < len(secretInt); i++ {
			// Checking if we correctly guessed the position and the value of a decimal
			if secretInt[i] == enteredValue[i] {
				positionAndValue++
				// If we did not guess both of them correctly we check if we have that decimal at all in the game
			} else if _, ok := decimalMap[int(enteredValue[i])-48]; ok {
				// int(enteredValue[i])-48 this is just turning an ascii value of an int to the real int
				justValue++
			}
		}
		fmt.Printf("There are %d decimals that you correctly guessed both the position and the value \n", positionAndValue)
		fmt.Printf("There are %d decimals that you correctly guessed only the value \n", justValue)
		fmt.Scanln(&enteredValue)
	}
	fmt.Println("Correct Guess Good Job")
}
func createRandomInt() (string, map[int]int) {
	decimal := rand.Intn(10)
	number := 0
	decimalPoint := 1
	var decimalMap = map[int]int{}
	for number < 1000 {
		if _, ok := decimalMap[decimal]; !ok {
			// Checking if the leading decimal is zero if so getting a new int
			if decimalPoint == 1000 && decimal == 0 {
				decimal = rand.Intn(10)
				continue
			}
			decimalMap[decimal] = 1
			number += decimal * decimalPoint
			decimalPoint *= 10
			decimal = rand.Intn(10)
		} else {
			decimal = rand.Intn(10)
		}
	}
	fmt.Println(number)
	return strconv.Itoa(number), decimalMap
}
