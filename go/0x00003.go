package main

import "fmt"

var a string

func main() {
	a = "G"
	print(a)
	f1()
	fmt.Println()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
