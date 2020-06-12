package main

import "fmt"

func main() {

	type person struct {
		fullName string
		age      int
	}

	type student struct {
		campus string
		person
	}

	var ramdhan student
	ramdhan.age = 22
	ramdhan.fullName = "Mochammad Ramadhan Januar Hidayat"
	ramdhan.campus = "widyatama"

	fmt.Println("Fullname:", ramdhan.fullName)
	fmt.Println("Age:", ramdhan.age)
	fmt.Println("Campus:", ramdhan.campus)

	var ramdhanCopy *student
	ramdhanCopy = &ramdhan
	ramdhanCopy.fullName = "hehehe"

	fmt.Println("Fullname:", ramdhan.fullName)
	fmt.Println("Fullname:", ramdhanCopy.fullName)
	fmt.Println("Age:", ramdhanCopy.age)
	fmt.Println("Campus:", ramdhanCopy.campus)

	var pujiPerson = person{fullName: "Puji Dwi Wahyuni", age: 22}
	var puji = student{person: pujiPerson, campus: "Unpad"}

	fmt.Println("Fullname:", puji.fullName)
	fmt.Println("Age:", puji.age)
	fmt.Println("Campus:", puji.campus)

	// Anonymous Struct

	var s1 = struct {
		person
		campus string
	}{
		person: pujiPerson,
		campus: "Unpad",
	}

	fmt.Println("Fullname:", s1.fullName)
	fmt.Println("Age:", s1.age)
	fmt.Println("Campus:", s1.campus)

	// slice struct
	var allStudent = []person{
		{fullName: "ramadhan", age: 22},
		{fullName: "puji", age: 22},
	}

	for _, student := range allStudent {
		fmt.Println("Fullname: ", student.fullName, "Age: ", student.age)
	}

	//anonymous slice struct

	var allStudentTwo = []struct {
		fullName string
		age      int
	}{
		{fullName: "ramadhan", age: 22},
		{fullName: "puji", age: 22},
	}

	for _, student := range allStudentTwo {
		fmt.Println("Fullname2: ", student.fullName, "Age: ", student.age)
	}

	var studentanon struct {
		person
		grade int
	}

	studentanon.person = person{fullName: "ramadhan", age: 22}
	studentanon.grade = 100

	fmt.Println("Fullname2: ", studentanon.fullName, "Age: ", studentanon.grade)
}
