package main

import (
	"fmt"
)

func main() {

	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a == b && a == c {
		fmt.Println("1")
	} else if (a == b && a != c) || (a != b && a == c) {
		fmt.Println("2")
	} else if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("0")
	} else {
		fmt.Println("3")
	}
}
