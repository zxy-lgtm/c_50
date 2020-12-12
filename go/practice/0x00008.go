package main

import "fmt"

func fp(a *[3]int) { fmt.Println(a) }
func fp2(a [3]int) {
	for i := 0; i < 3; i++ {
		fmt.Print(" ", a[i])
	}
	fmt.Println()
}
func fp3(a [3]int) {
	fmt.Println(a)
}

func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
		fp2([3]int{i, i + i, i + i + i})
		fp3([3]int{i - 1, i + i - 1, i + i + i - 1})
	}

	ch := [2][2]string{{"3", "we"}, {"4", "you"}}
	var ch1 [2][2]string

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			ch1[i][j] = "we"
		}
	}

	fmt.Println()
	fmt.Println(ch)
	fmt.Println(ch1)

}
