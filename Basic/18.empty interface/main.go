package main

import (
	"fmt"
	"strings"
)

func main() {
	var everything interface{}

	everything = "halo ramdhan"
	fmt.Println(everything)

	everything = 10.10
	fmt.Println(everything)

	everything = []string{"ramdhan", "puji"}
	fmt.Println(everything)

	everything = map[string]interface{}{
		"name": "Ramdhan",
		"age":  22,
	}

	fmt.Println(everything)

	data := map[string]interface{}{
		"name": "ramdhan",
		"age":  22,
	}

	fmt.Println(data["name"])
	fmt.Println(data["age"])

	//need to casting if u using interface{} data type
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	type person struct {
		name string
		age  int
	}

	var secret2 interface{} = &person{name: "wick", age: 27}
	var name = secret2.(*person).name
	fmt.Println(name)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	var persons = []map[string]interface{}{
		{"name": "Ramdhan", "age": 22},
		{"name": "Ramdhan", "age": 22},
		{"name": "Ramdhan", "age": 22},
		{"name": "Ramdhan", "age": 22},
	}

	for _, person := range persons {
		fmt.Println(person["name"], person["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
