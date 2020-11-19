package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][0] = 2
	}
	fmt.Println(items)

	items[1][2] = 3
	fmt.Println(items)

	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)

	}

	fmt.Println("\n")

	keys := make([]string, len(barVal))
	count := 0
	for i, _ := range barVal {
		keys[count] = i
		count++
	}

	fmt.Println()

	for i, j := range barVal {
		fmt.Println(i, j)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
