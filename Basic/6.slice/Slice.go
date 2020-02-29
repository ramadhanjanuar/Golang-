package main

import "fmt"

func main() {
	var fruitsA = []string{"apple", "grape"}   // slice
	var fruitsB = [2]string{"banana", "melon"} // array
	var newFruits = fruitsA[0:2]
	var newFruitsB = newFruits[0:2]

	for _, fruit := range fruitsA {
		fmt.Println(fruit)
	}

	for _, fruit := range fruitsB {
		fmt.Println(fruit)
	}

	newFruitsB[0] = "mantol"

	for _, fruit := range newFruits {
		fmt.Println(fruit)
	}

	for _, fruit := range newFruitsB {
		fmt.Println(fruit)
	}

	var aaFruits = []string{
		"apple",
		"banana",
		"pinapple",
		"strawberry",
	}

	var aaaFruits = aaFruits[0:3]

	fmt.Println(len(aaFruits))
	fmt.Println(cap(aaFruits))

	fmt.Println(len(aaaFruits))
	fmt.Println(cap(aaaFruits))

}
