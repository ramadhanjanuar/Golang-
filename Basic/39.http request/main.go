package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type student struct {
	ID   string
	Name string
	Age  int
}

var baseURL = "http://localhost:8080"

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("ID", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var users, err = fetchUsers()

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, user := range users {
		fmt.Println("ID   :", user.ID)
		fmt.Println("Name :", user.Name)
		fmt.Println("Age  :", user.Age)
		fmt.Println("")
	}

	user1, err := fetchUser("E001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Age: %d\n", user1.ID, user1.Name, user1.Age)
}
