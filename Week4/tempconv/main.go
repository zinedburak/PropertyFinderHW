package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"tempconv/conversions"
)

// Convertor struct that has different temperatures types
type Convertor struct {
	tempTypes []conversions.Conversion
}

// addType is for adding different types to convertor
func (c *Convertor) addType(tempType conversions.Conversion) {
	c.tempTypes = append(c.tempTypes, tempType)
}

// convert is the main function that finds which and converts from A type to B type
func (c Convertor) convert(typeName string, value float64, convertType string) (float64, error) {
	typeName = strings.ToLower(typeName)
	for _, tempType := range c.tempTypes {
		if strings.Contains(strings.ToLower(reflect.TypeOf(tempType).String()), typeName) {
			converted, _, err := tempType.Convert(value, convertType)
			if err != nil {
				return 0, err
			}
			return converted, nil
		}
	}
	return 0, errors.New("the typename that you inputted is not one of the types that are available")
}

func main() {

	initConvertor()

}

func initConvertor() {
	convertor := Convertor{}

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("\nConvertor started.")
	fmt.Println("You can add the fallowing types anc convert them to the other types in the list")

	convertor.addType(conversions.Kelvin(0))
	convertor.addType(conversions.Celsius(0))
	convertor.addType(conversions.Fahrenheit(0))

	for _, v := range convertor.tempTypes {
		typeNames := strings.Replace(reflect.TypeOf(v).String(), "conversions.", "", 1)
		fmt.Println(typeNames)
	}

	for input.Scan() {

		var typeName string
		var value float64
		var convertType string

		fmt.Println("> Enter typename value and convert type with one space")
		_, err := fmt.Scanf("%s %f %s", &typeName, &value, &convertType)
		if err != nil {
			fmt.Println("Wrong input type in one of the input values please check Error: ", err)
			continue
		}
		if typeName == "x" {
			os.Exit(0)
		}
		converted, err := convertor.convert(typeName, value, convertType)
		if err == nil {
			fmt.Printf("%f %s is %f %s", value, typeName, converted, convertType)
		} else {
			fmt.Println(err)
			continue
		}

	}

}
