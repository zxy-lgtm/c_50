package main

import "fmt"

func main() {
	a := 2
	b := 1
	fmt.Println(a)
	fmt.Println(a|b, a&b, a^b)
	a = b << 10
	b = a >> 5
	fmt.Println(a, b)
}
