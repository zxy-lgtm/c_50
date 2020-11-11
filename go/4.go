package main

import (
	"fmt"
)

func main() {
	var number1, number2, divisor int
	fmt.Scanf("%d %d", &number1, &number2)
	divisor = common_divisor(number1, number2)
	fmt.Println(divisor)
}

func common_divisor(n1 int, n2 int) int {

	for n1 != n2 {
		if n1 > n2 {
			n1 = n1 - n2
		} else {
			n2 = n2 - n1
		}
	}
	return n1

}
