package main

import "fmt"

func main() {
	//Initialization map
	var day = map[string]string{}

	//Assign value to map
	day["one"] = "monday"
	day["two"] = "tuesday"
	day["three"] = "wednesday"
	day["four"] = "thursday"
	day["five"] = "friday"
	day["six"] = "saturday"
	day["seven"] = "sunday"
	fmt.Println(day)

	//Iteration Item
	for key, val := range day {
		fmt.Println(key, " :", val)
	}

	//Delete Item
	delete(day, "one")
	fmt.Println(day)

	//check exist value in map
	var value, isExist = day["four"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("value is not exist")
	}

	//map slice
	var persons = []map[string]string{
		{"name": "ramdhan", "gender": "male"},
		{"name": "puji", "gender": "female"},
	}

	for _, person := range persons {
		fmt.Println("my name:", person["name"], "| my gender:", person["gender"])
	}

}
