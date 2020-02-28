package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var i uint = 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	fmt.Println('\n')

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 19 {
			break
		}
		fmt.Println(i)
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
