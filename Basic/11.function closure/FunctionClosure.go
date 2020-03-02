package main

import "fmt"

func main() {
	var getMinMax = func(numbers []int) (min, max int) {
		for i, number := range numbers {
			switch {
			case i == 0:
				min, max = number, number
			case number > max:
				max = number
			case number < min:
				min = number
			}
		}

		return
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	var newNumbers = func(min int) []int {
		var result []int
		for _, number := range numbers {
			if number < min {
				continue
			}
			result = append(result, number)
		}
		return result
	}(3)

	fmt.Println("original number : ", numbers)
	fmt.Println("filtered number : ", newNumbers)
}
