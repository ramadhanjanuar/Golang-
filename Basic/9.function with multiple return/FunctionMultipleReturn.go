package main

import (
	"fmt"
	"math"
)

func main() {
	var area, circumference = calculateOne(10)
	fmt.Println(area, circumference)

	var areaTwo, circumferenceTwo = calculateOne(10)
	fmt.Println(areaTwo, circumferenceTwo)
}

func calculateOne(d float64) (float64, float64) {
	// calculate area
	var area = math.Pi * math.Pow(d/2, 2)

	// calculate circumference
	var circumference = math.Pi * d

	return area, circumference
}

func calculateTwo(d float64) (area, circumference float64) {
	// calculate area
	area = math.Pi * math.Pow(d/2, 2)

	// calculate circumference
	circumference = math.Pi * d

	return
}
