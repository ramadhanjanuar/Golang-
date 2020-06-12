package main

import "fmt"

const multiplyResult int = 20

var data = []int{2, 4, 1, 5, 6, 40, -1}

// TwoNumbers of two numbers
type TwoNumbers struct {
	NumberOne int
	NumberTwo int
}

// ThreeNumbers of pair numbers
type ThreeNumbers struct {
	NumberOne   int
	NumberTwo   int
	NumberThree int
}

func main() {
	resultTwoNumbers := findTwoIntegerByMultiplyResult(multiplyResult)
	resultThreeNumbers := findThreeIntegerByMultiplyResult(multiplyResult)

	fmt.Printf("List of Pair Numbers: %v \n", resultTwoNumbers)
	fmt.Printf("List of Pair Numbers: %v \n", resultThreeNumbers)
}

func findTwoIntegerByMultiplyResult(result int) []TwoNumbers {
	var listOfResult []TwoNumbers
	for _, number1 := range data {
		for _, number2 := range data {
			if (number1 * number2) == result {
				pairNumbers := TwoNumbers{NumberOne: number1, NumberTwo: number2}
				listOfResult = append(listOfResult, pairNumbers)
			}
		}
	}
	return listOfResult
}

func findThreeIntegerByMultiplyResult(result int) []ThreeNumbers {
	var listOfResult []ThreeNumbers
	for _, number1 := range data {
		for _, number2 := range data {
			for _, number3 := range data {
				if (number1 * number2 * number3) == result {
					ThreeNumbers := ThreeNumbers{NumberOne: number1, NumberTwo: number2, NumberThree: number3}
					listOfResult = append(listOfResult, ThreeNumbers)
				}
			}
		}
	}
	return listOfResult
}
