package main

import (
	"fmt"
	"math"
)

type Cube struct {
	Side float64
}

func (k Cube) Volume() float64 {
	return math.Pow(k.Side, 3)
}

func (k Cube) Area() float64 {
	return math.Pow(k.Side, 2) * 6
}

func (k Cube) Around() float64 {
	return k.Side * 12
}

func main() {
	fmt.Println("Hello World")
}
