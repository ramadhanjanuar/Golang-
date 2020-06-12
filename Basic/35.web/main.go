package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo")
	})

	http.HandleFunc("/index", index)

	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Ramadhan",
			"Message": "Hello, Welcom to our web:)",
		}

		var t, err = template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("Web runnin at 8080")
	http.ListenAndServe(":8080", nil)
}
