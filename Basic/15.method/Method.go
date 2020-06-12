package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	age  int
}

func (s student) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s student) getNameAt(number int) string {
	if lengthName := len(strings.Split(s.name, " ")); lengthName <= 1 {
		return s.name
	}
	return strings.Split(s.name, " ")[number-1]
}

func main() {
	var s1 = student{
		name: "Ramdhan",
		age:  22,
	}

	s1.sayHello()
	s1.getNameAt(2)

	s1.changeName("puji")
	fmt.Println(s1.name)
	s1.changeNamePointer("puji")
	fmt.Println(s1.name)
}

func (s student) changeName(newName string) {
	s.name = newName
}

func (s *student) changeNamePointer(newName string) {
	s.name = newName
}
