package main

import "fmt"

func main() {
	var numbers = []int{2, 4, 1, 5, 3, 6, 7, 8, 10}
	var avg = average(10, 2, 3, 4, 1, 4, 6, 7)
	var msg = fmt.Sprintf("Rata-rata %.2f", avg)
	fmt.Println(msg)

	var avgTwo = average(numbers...)
	var msgTwo = fmt.Sprintf("Rata-rata %.2f", avgTwo)
	fmt.Println(msgTwo)
}

func average(numbers ...int) (average float64) {
	var total = 0

	for _, number := range numbers {
		total += number
	}

	average = float64(total) / float64(len(numbers))

	return
}
