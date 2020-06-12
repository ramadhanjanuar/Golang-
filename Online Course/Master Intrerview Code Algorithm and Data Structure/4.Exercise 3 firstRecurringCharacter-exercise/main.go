//Google Question
//Given an array = [2,5,1,2,3,5,1,2,4]:
//It should return 2

//Given an array = [2,1,1,2,3,5,1,2,4]:
//It should return 1

//Given an array = [2,3,4,5]:
//It should return undefined

//Bonus... What if we had this:
// [2,5,5,2,3,5,1,2,4]
// return 5 because the pairs are before 2,2

package main

import "fmt"

func main() {
	data1 := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}

	findFirstRecurringCharInSlice(data1)
	x := findFirstRecurringCharInSliceEfficien(data1)
	fmt.Println(x)
}

// O(n^2)
func findFirstRecurringCharInSlice(array []int) {
	var tempChar map[int]int
	tempChar = map[int]int{}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			var _, isExist = tempChar[array[i]]
			if array[i] == array[j] && !isExist {
				tempChar[array[i]] = j
			}
		}
	}

	var min int
	var result int
	for key, value := range tempChar {
		if min == 0 {
			min = value
			result = key
		}
		if value < min {
			min = value
			result = key
		}
	}
	if result == 0 && min == 0 {
		fmt.Println("undefined")
	}
	fmt.Println(result)
}

// O(N)
func findFirstRecurringCharInSliceEfficien(array []int) int {
	var mapVar map[int]int
	mapVar = map[int]int{}

	for i := 0; i < len(array); i++ {
		_, isExist := mapVar[array[i]]
		if isExist == true {
			return array[i]
		}
		mapVar[array[i]] = array[i]
	}
	return 0
}
