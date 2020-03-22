package main

import (
	"encoding/json"
	"fmt"
)

// User is struct for user
type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "Mochammad Ramadhan Januar", "Age": 22}`
	var jsonData = []byte(jsonString)

	var user User

	var err = json.Unmarshal(jsonData, &user)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Fullname:", user.Fullname)
	fmt.Println("Age:", user.Age)

	var userMap map[string]interface{}
	json.Unmarshal(jsonData, &userMap)

	fmt.Println("Fullname:", userMap["Name"])
	fmt.Println("Age:", userMap["Age"])

	var jsonUsersString = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var jsonUsersData = []byte(jsonUsersString)

	var users []User
	err = json.Unmarshal(jsonUsersData, &users)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Fullname:", users[0].Fullname)
	fmt.Println("Age:", users[0].Age)

	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonDataFromObject, err2 = json.Marshal(object)

	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	var jsonStringFromObject = string(jsonDataFromObject)
	fmt.Println(jsonStringFromObject)

}
