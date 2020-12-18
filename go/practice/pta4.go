package main

import (
	"fmt"
)

type Student struct {
	Name string
	ID   string
}

func main() {
	students := make(map[int]Student)
	var n int
	min := 100
	max := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var name, id string
		var score int
		var student Student
		fmt.Scan(&name, &id, &score)
		student.Name = name
		student.ID = id
		students[score] = student
		if score < min {
			min = score
		}

		if score > max {
			max = score
		}
		//fmt.Println(students)
	}
	studentmax := students[max]
	studentmin := students[min]
	fmt.Println(studentmax.Name, studentmax.ID)
	fmt.Println(studentmin.Name, studentmin.ID)

}
