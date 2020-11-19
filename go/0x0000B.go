package main

import "fmt"

const (
	Len = 10
)

type fibonacci []int

func Fibonacci(Len int) fibonacci {
	fibonacci := make([]int, Len)
	fibonacci[0] = 1
	fibonacci[1] = 1
	for i := 2; i < Len; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	return fibonacci

}

func main() {
	slice1 := make([]string, 5, 10)
	slice1[2] = "5"
	slice2 := new([10]int)[2:8]
	slice2[4] = 9
	fmt.Println(slice1)
	fmt.Println(slice2)
	fibonacci := Fibonacci(Len)
	fmt.Println(fibonacci)

}
