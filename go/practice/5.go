package main

import "fmt"

func main() {

	count := 0

	for j := 2; j < 101; j++ {
		error := prime(j)
		if error {
			print(" ", j)
			count++

			if count%5 == 0 {
				fmt.Println()
			}
		}
	}

}

func prime(number int) bool {

	count := 0

	for i := 2; i < number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 0 {
		return true
	} else {
		return false
	}
}
