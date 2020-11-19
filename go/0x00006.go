package main

import (
	"fmt"
)

func main() {
	callback(2, Add)
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron", 2: "Tim"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	for i := 0; i < 10; i++ {
		f(y, i*6)
	}
}
