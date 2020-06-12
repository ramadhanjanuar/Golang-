package main

import "fmt"

func main() {
	var array1 = []string{"a", "b", "c", "x"}
	var array2 = []string{"z", "c", "i"}
	result := twoArrayContains(array1, array2)
	result2 := twoArrayContainsBetter(array1, array2)
	fmt.Println("Result is", result)
	fmt.Println("Result is", result2)
}

// O(n^2) or O(a*b) if the size of array different
// O(1) space complexity
func twoArrayContains(array1, array2 []string) bool {
	for _, item := range array1 {
		for _, item2 := range array2 {
			if item == item2 {
				return true
			}
		}
	}
	return false
}

// O(n) or O(a+b)
// O(a) space complexity
func twoArrayContainsBetter(array1, array2 []string) bool {
	var obj = map[string]bool{}
	for _, item1 := range array1 {
		obj[string(item1)] = true
	}
	for _, item2 := range array2 {
		_, isExist := obj[item2]
		if isExist {
			return isExist
		}
	}
	return false
}
