package main

import "fmt"

func getAverage(data []int, ch chan float64) {
	var sum int
	var avg float64

	for _, number := range data {
		sum += number
	}

	avg = float64(sum) / float64(len(data))

	ch <- avg
}

func getMax(data []int, ch chan int) {
	var max int = 0

	for _, number := range data {
		if max < number {
			max = number
		}
	}

	ch <- max
}

func main() {
	var data = []int{10, 3, 4, 1, 3, 6, 4, 6, 9}

	var ch1 = make(chan float64)
	go getAverage(data, ch1)

	var ch2 = make(chan int)
	go getMax(data, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Println("Average :", avg)
		case max := <-ch2:
			fmt.Println("Max :", max)
		}
	}
}
