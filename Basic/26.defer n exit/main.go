package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hallo")
	defer fmt.Println("test aja")
	fmt.Println("Selamat Datang")
	orderSomeFood("pizza")
	orderSomeFood("burger")
	os.Exit(1)
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		defer fmt.Println("halo 3")
	}

	fmt.Println("halo 2")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
