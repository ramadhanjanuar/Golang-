package main

import (
	"fmt"
	"time"
)

func main() {
	//var nemo = []string{"nemo", "test", "apa", "aja", "gitu"}
	var slices = createSliceString(1000000, "nemo")
	findNemo(slices)
}

func createSliceString(number int, value string) []string {
	var slices []string
	for i := 0; i < number; i++ {
		slices = append(slices, value)
	}
	return slices
}

func findNemo(arrays []string) {
	strat := time.Now()
	for _, item := range arrays {
		if item == "nemo" {
			fmt.Println("Found Nemo")
			break
		}
	}
	duration := time.Since(strat)
	fmt.Println("Duration: ", duration.Milliseconds())
}
