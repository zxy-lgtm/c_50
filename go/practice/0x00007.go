package main

import "fmt"

func main() {
	var a = [10]string{1: "w", "3", "e", "r"}

	b := [...]int{1, 2, 3, 4}

	var Fibonacci = [50]int{0: 1, 1: 1}
	for i := 0; i < 48; i++ {
		Fibonacci[i+2] = Fibonacci[i+1] + Fibonacci[i]
	}
	for i := 0; i < 50; i++ {
		fmt.Print(" ", Fibonacci[i])
	}

	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Println(a[i], b[i])
	}
}
