package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("请输入数组的元素个数和右移位数:")
	var number, bit int

	fmt.Scan(&number, &bit)
	if bit > number {
		num := make([]int, number)
		fmt.Println(num)

		return
	}
	//fmt.Printf("请输入%v个数(中间包含空格)：", number)
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	if err != nil {
		return
	}
	s := []byte(str)
	num := make([]int, number)
	i := 0
	for key, value := range s {
		if value != ' ' && key < (number*2-1) {
			num[i] = int(value - '0')
			i++
		}
	}
	//fmt.Println(num)

	num = append(num[number-bit:number], num[:number-bit]...)
	for key, value := range num {
		if key == 0 {
			fmt.Printf("%d", value)
		} else {
			fmt.Printf(" %d", value)
		}
	}
	fmt.Println()
}
