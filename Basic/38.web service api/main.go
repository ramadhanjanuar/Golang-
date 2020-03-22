package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID   string
	Name string
	Age  int
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"W001", "wick", 22},
	student{"B001", "bourne", 23},
	student{"B002", "bond", 23},
}

// HandleUsers is Handler for get Users
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method == "POST" {
		var jsonData, err = json.Marshal(data)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Write(jsonData)
	}

	http.Error(w, "status", http.StatusBadRequest)
}

// HandleUser is Handler for get User
func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("ID")

		for _, user := range data {
			if user.ID == id {
				var jsonData, err = json.Marshal(user)

				if err != nil {
					fmt.Println(err.Error())

				}
				w.Write(jsonData)
				return
			}
		}

		http.Error(w, "user not found", http.StatusBadRequest)
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", HandleUsers)
	http.HandleFunc("/user", HandleUser)

	fmt.Println("server running on 8080")
	http.ListenAndServe(":8080", nil)
}
