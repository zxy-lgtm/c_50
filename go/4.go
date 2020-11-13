package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	fmt.Scanf("%d %d", &number1, &number2)
	divisor1 := common_divisor1(number1, number2)
	fmt.Println(divisor1)
	divisor2 := common_divisor2(number1, number2)
	fmt.Println(divisor2)
}

func common_divisor1(n1 int, n2 int) int {

	for n1 != n2 {
		if n1 > n2 {
			n1 = n1 - n2
		} else {
			n2 = n2 - n1
		}
	}
	return n1
}

func common_divisor2(n1 int, n2 int) int {

	var divisor int
	if n1%n2 == 0 {
		divisor = n2
	} else {
		divisor = common_divisor2(n2, n1%n2)
	}
	return divisor
}
