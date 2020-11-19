package main

import "fmt"

const (
	WIDTH  = 9
	HEIGHT = 1
)

type pixel int

var sum int

var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 1
		}
	}

	fmt.Println(screen)
	var w = [3]int{2, 3, 4}
	var e = &w
	fp(e)
	fmt.Println(sum)
	var rt = [][]string{{"we", "r"}, {"3"}}
	rt[0][1] = "you"
	var r = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(r)
	r = r[:6]
	fmt.Println(r)
	r = r[4:7]
	fmt.Println(r)
	r = r[:cap(r)]
	fmt.Println(r)
	fmt.Println(rt)
}

func fp(a *[3]int) int {
	for i := 0; i < 3; i++ {
		sum += a[i]
	}
	return sum
}
