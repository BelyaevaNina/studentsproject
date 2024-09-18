package main

import "fmt"

type Student struct {
	name    string
	surname string
	grade   int
}

var students = []Student{}

func main() {

	var tom = Student{
		name:    "Tom",
		surname: "Person",
		grade:   3,
	}
	var grey = Student{
		name:    "Grey",
		surname: "Silver",
		grade:   4,
	}
	students = append(students, tom)
	students = append(students, grey)
	fmt.Println(students)
}
