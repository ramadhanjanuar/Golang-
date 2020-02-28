package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = []string{"Ramadhan", "Januar"}
	printMessage("hallo", name)
	fmt.Println(1, " is ", evenOdd(8))
	devidNumber(10, 2)
}

func printMessage(message string, arr []string) {
	var arrayString = strings.Join(arr, " ")
	fmt.Println(message, arrayString)
}

func evenOdd(number int) string {
	if number%2 == 0 {
		return "even"
	}

	return "odd"
}

func devidNumber(valA, valB int) {
	if valB == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", valA, valB)
	}

	var result = valA / valB
	fmt.Printf("%d / %d = %d ", valA, valB, result)

}
