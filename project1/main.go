package main

import "fmt"

type Student struct {
	Name   string
	Number string
	Age    int
}

func main() {
	fmt.Println("Hello World")

    students:= map[string]Student {
		"45885" : {"Ali","45885",18},
		"74832" : {"Amir","74832",21},
		"23422" : {"Shahrzad","23422",17},
	}

	var student_number string
	fmt.Printf("Enter student number : ")
	fmt.Scan(&student_number)

	student, ok := students[student_number]
	if ok {
		fmt.Printf("Name : %s  Student number : %v   age : %v ",student.Name,student.Number,student.Age)
	} else {
		fmt.Println("student not found, make sure you've entered the correct student number")
	}
}
