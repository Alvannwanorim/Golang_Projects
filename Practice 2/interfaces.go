package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type rect struct {
	Height float64
	Width  float64
}

type circle struct {
	Radius float64
}

func (r rect) area() float64 {
	return r.Height * r.Width
}

func (c circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func checkType(s shape) {
	c, ok := s.(circle)
	if ok {
		fmt.Println("This is a circle", c)
	}
	c1, ok := s.(rect)
	if ok {
		fmt.Println("This is a rectangle", c1)
	}
}
