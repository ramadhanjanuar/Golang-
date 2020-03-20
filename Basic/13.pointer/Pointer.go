package main

import "fmt"

func main() {
	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println("Number A Value :", numberA)
	fmt.Println("Number A MemLoc :", &numberA)

	fmt.Println("Number B Value :", *numberB)
	fmt.Println("Number B MemLoc :", numberB)

	numberA = 10

	fmt.Println("Number A Value :", numberA)
	fmt.Println("Number B Value :", *numberB)

	changeValueVariablePointer(numberB, 15)

	fmt.Println("Number A Value :", numberA)
	fmt.Println("Number B Value :", *numberB)
}

func changeValueVariablePointer(pointer *int, value int) {
	*pointer = value
}
