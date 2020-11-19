package main

import (
	"fmt"
	"strings"
)

type Name struct {
	fristName string
	lastName  string
}

func upName(ptr *Name) {
	ptr.fristName = strings.ToUpper(ptr.fristName)
	ptr.lastName = strings.ToUpper(ptr.lastName)
}
func main() {
	var person1 Name
	person1.fristName = "li"
	person1.lastName = "jie"
	ptr1 := &person1
	upName(ptr1)
	fmt.Println(person1)

	person2 := Name{"wang", "siying"}
	upName(&person2)
	fmt.Println(person2)

	person3 := new(Name)
	person3.fristName = "zhang"
	person3.lastName = "die"
	(*person3).lastName = "die"
	upName(person3)
	fmt.Println(person3)

}
