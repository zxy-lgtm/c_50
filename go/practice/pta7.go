package main

import "fmt"

func PrimeNumber(n int) int {
	flag := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag++
		}
	}
	if flag == 0 {
		return n
	}
	return 0
}

func main() {
	var n int
	tool := 0
	result := 0
	var prime [10000]int
	fmt.Scan(&n)
	//fmt.Println(n)
	for i := 2; i <= n; i++ {
		if PrimeNumber(i) != 0 {
			prime[tool] = PrimeNumber(i)
			tool++
			//fmt.Println(PrimeNumber(i))
		}
	}
	//fmt.Println(prime, tool)

	for i := 0; i < tool; i++ {
		if prime[i+1]-prime[i] == 2 {
			result++
		}
	}

	fmt.Println(result)
}
