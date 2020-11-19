package main

import "fmt"

func main() {
	var sum int
	var min int
	items := [...]int{10, 20, 30, 1, 89, 33, 40, 50}
	for _, item := range items {
		item *= 2
		fmt.Println(item)
	}
	for _, number := range items {
		sum += number
	}
	min = items[0]
	for _, number := range items {
		if number < min {
			min = number
		}
	}
	fmt.Println(items, sum, min)
}
