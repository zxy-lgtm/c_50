package main

import "fmt"

//条件是int32高位溢出
func reverse(x int) int {
	var number int
	for x != 0 {
		number = number*10 + x%10
		if !(-(1<<31) <= number && number <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return number
}

func main() {
	x := 123
	x = reverse(x)
	fmt.Println("反转后的整数是:")
	fmt.Println(x)
}
