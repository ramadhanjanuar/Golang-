package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getInfoAllProperty() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("name : ", reflectType.Field(i).Name)
		fmt.Println("type :", reflectType.Field(i).Type)
		fmt.Println("Nilai:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 5
	var reflectValue = reflect.ValueOf(number)

	fmt.Println(reflectValue)
	fmt.Println(reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println(reflectValue.Int())
	}
	fmt.Println(reflectValue.Interface())

	var s1 = &student{Name: "Ramdhan", Grade: 100}
	reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println(s1.Name)
}
