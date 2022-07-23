package conversions

import (
	"errors"
	"strings"
)

const (
	fahrenheit = "fahrenheit"
	celsius    = "celsius"
	kelvin     = "kelvin"
)

type Conversion interface {
	Convert(value float64, conversion string) (float64, string, error)
}

type Kelvin float64

// Conversion functions of kelvin it takes value and the type of what it should convert that value

func (k Kelvin) Convert(value float64, conversion string) (float64, string, error) {
	conversion = strings.ToLower(conversion)
	if conversion == strings.ToLower(celsius) {
		// celsius conversion from kelvin
		return value - 273.15, "Celsius", nil
	} else if conversion == strings.ToLower(fahrenheit) {
		// fahrenheit conversion from kelvin
		return (value-273.15)*9/5 + 32, "Fahrenheit", nil
	} else {
		return 0, "", errors.New("there is no conversion type from Kelvin to your input")
	}

}

type Celsius float64

// Conversion functions of Celsius it takes value and the type of what it should convert that value

func (c Celsius) Convert(value float64, conversion string) (float64, string, error) {
	conversion = strings.ToLower(conversion)
	if conversion == strings.ToLower(fahrenheit) {
		//  fahrenheit conversion from celsius
		return value*9/5 + 32, "Fahrenheit", nil
	} else if conversion == strings.ToLower(kelvin) {
		//  kelvin conversion from celsius
		return value + 273.15, "Kelvin", nil
	} else {
		return 0, "", errors.New("there is no conversion type from Celsius to your input")
	}
}

type Fahrenheit float64

// Conversion functions of Fahrenheit it takes value and the type of what it should convert that value

func (f Fahrenheit) Convert(value float64, conversion string) (float64, string, error) {
	conversion = strings.ToLower(conversion)
	if conversion == strings.ToLower(celsius) {
		//  celsius conversion from fahrenheit
		return (value - 32) * 5 / 9, "", nil
	} else if conversion == strings.ToLower(kelvin) {
		//  kelvin conversion from fahrenheit
		return (value-32)*5/9 + 273.15, "", nil
	} else {
		return 0, "", errors.New("there is no conversion type from Fahrenheit to your input")
	}
}
