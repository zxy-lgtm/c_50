package main

import "fmt"

func Divisor(n int) (int, int) {
	div := n % 10
	number := n / 10
	return div, number
}

func main() {
	var n int
	fmt.Scan(&n)
	number := 1
	//fmt.Println(n)
	div1, n := Divisor(n)
	div2, div3 := Divisor(n)
	//div1,div2,div3分别是个位数,十位数,百位数
	//fmt.Println(div1, div2, div3)
	for ; div3 > 0; div3-- {
		fmt.Print("B")
	}
	for ; div2 > 0; div2-- {
		fmt.Print("S")
	}
	for ; div1 > 0; div1-- {

		fmt.Print(number)
		number++
	}
	fmt.Println()
}
