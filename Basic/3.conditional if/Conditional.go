package main

import "fmt"

func main() {
	const luckyNumber uint = 6

	if luckyNumber == 6 {
		fmt.Println("ur lucky number")
	} else if luckyNumber >= 3 && luckyNumber <= 6 {
		fmt.Println("almost get lucky number")
	} else {
		fmt.Println("wrong number")
	}

	if myLuckyNumber := luckyNumber / 2; myLuckyNumber == 3 {
		fmt.Println("ur my lucky number")
	} else if myLuckyNumber == 1 {
		fmt.Println("So WORNG!!!!")
	} else {
		fmt.Println("damn!!!", myLuckyNumber)
	}

	switch luckyNumber {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 6:
		fmt.Println("MSANTOL")
	default:
		fmt.Println("UR SO LUCKY!!!1")
	}

	switch {
	case luckyNumber == 5:
		fmt.Println("mantol")
	case (luckyNumber > 3) && (luckyNumber < 7):
		fmt.Println("apa yang terjadi hehe")
	default:
		fmt.Println("default nya ini")
	}
}
