package main

import (
	"fmt"
	"runtime"
)

func print(till int, word string) {
	for i := 0; i < till; i++ {
		fmt.Println(word)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	var names [4]string

	for i := range names {
		fmt.Scan(&input)
		names[i] = input
	}

	fmt.Println(names)
}
