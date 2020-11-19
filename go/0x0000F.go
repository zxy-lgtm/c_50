package main

import "fmt"

func main() {
	map1 := map[string]int{"one": 1, "two": 2}
	map1["one"] = 4
	fmt.Println(map1["one"], map1["ten"], map1) //map1["ten"] == 0
	map2 := map[int]float64{1: 2.3, 2: 3.4}
	fmt.Println(map2[1], map2[10]) //map2[10] == 0
	map3 := make(map[string]int)
	map3["you"] = 1
	fmt.Println(map3["you"], map3["I"])
	map4 := map[int]func() int{
		1: func() int { return 1 },
	}
	fmt.Println(map4[1])

	map5 := make(map[int]float64, 100)
	map5[1] = 2.3
	fmt.Println(map5)

	//判断map5中的值是否存在
	_, ok1 := map5[1]
	_, ok2 := map5[2]
	fmt.Printf("%t %t", ok1, ok2)
	fmt.Println()

	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	fmt.Println(noteFrequency) //输出了map[A0:27.5 A4:440 B0:30.87 C0:16.35 D0:18.35 E0:20.6 F0:21.83 G0:24.5]

	map6 := make(map[int][]int)
	number := []int{1, 2, 34, 5, 6, 7}
	slice1 := number[:]
	map6[1] = slice1
	fmt.Println(map6) //map[1:[1 2 34 5 6 7]]

	map7 := map[int]int{1: 2, 2: 3, 3: 4, 4: 5}
	for a, b := range map7 {
		if a+b == 5 {
			fmt.Println(a, "+", b, "=", a+b)
		}
	}

	map8 := map[int]float64{1: 2.3, 2: 3.4, 8: 2.3, 3: 2.3}
	map9 := make(map[int]int)
	map9[1] = 8
	map9[2] = 5
	map9[3] = 3
	map9[4] = 1
	for a, b := range map8 {
		fmt.Println(a, b)
	}
	for a, b := range map9 {
		fmt.Println(a, b)
	}

}
