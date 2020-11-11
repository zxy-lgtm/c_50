package main

import (
	"fmt"
	"strconv"
)

func main() {

	s1 := "777"
	fmt.Println(s1)
	a, _ := strconv.Atoi(s1)
	fmt.Println(a)
	a += 5
	fmt.Println(a)
	s2 := strconv.Itoa(a)
	fmt.Println(s2)
	fmt.Println(strconv.IntSize)
}
