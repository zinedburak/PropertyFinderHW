package models

import "fmt"

type Rectangle struct {
	ShortEdge float64
	LongEdge  float64
}

func (r Rectangle) Area() float64 {
	return r.ShortEdge * r.LongEdge
}
func (r Rectangle) Circumference() float64 {
	return 2 * (r.ShortEdge + r.LongEdge)
}
func CreateRectangle(shortEdge float64, longEdge float64) (Rectangle, error) {
	if shortEdge <= 0 {
		return Rectangle{}, fmt.Errorf("short edge cannot be smaller than zero")
	}
	if longEdge <= 0 {
		return Rectangle{}, fmt.Errorf("long edge cannot be smaller than zero")
	}
	if longEdge < shortEdge {
		return Rectangle{}, fmt.Errorf("short edge cannot be larger than long edge")
	}
	return Rectangle{shortEdge, longEdge}, nil
}
