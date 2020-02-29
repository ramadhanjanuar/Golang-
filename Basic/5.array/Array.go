package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Ramadhan"
	names[1] = "Puji"

	fmt.Println(names[0], names[1])

	var loves = [3]string{
		"aku",
		"sayanag",
		"kamu",
	}

	fmt.Println(loves)

	var numbers = [...]int{
		1,
		2,
		3,
	}

	fmt.Println(numbers)

	var numbersOne = [2][4]int{[4]int{1, 2, 3, 4}, [4]int{2, 3, 4, 5}}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbersOne[1][i])
	}

	for _, love := range loves {
		fmt.Println(love)
	}

	var fruits = make([]string, 2)
	fruits[0] = "test"
	fruits[1] = "test2"
	fmt.Println(fruits)
}
