package main

import (
	"fmt"
	"log"
	"rectangle/models"
)

func main() {
	rectangle0 := models.Rectangle{
		ShortEdge: 10,
		LongEdge:  100,
	}
	fmt.Println("Area of rectangle 0 : ", rectangle0.Area())
	fmt.Println("Circumference of rectangle 0 : ", rectangle0.Circumference())

	rectangle1, err := models.CreateRectangle(10.0, 100.5)
	if err != nil {
		log.Println("Error Creating rectangle : ", err)
	}
	fmt.Println("Area of rectangle 0 : ", rectangle1.Area())
	fmt.Println("Circumference of rectangle 0 : ", rectangle1.Circumference())

	_, err = models.CreateRectangle(5, 2) // Rectangle with invalid edges
	if err != nil {
		log.Println("Error Creating rectangle : ", err)
	}

	_, err = models.CreateRectangle(-5, -10) // Rectangle with edges that are smaller than zero
	if err != nil {
		log.Println("Error Creating rectangle : ", err)
	}
}
