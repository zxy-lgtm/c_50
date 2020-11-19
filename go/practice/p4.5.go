package main

import "fmt"

type Rope string

var ch Rope

func main() {
	ch = "lemon"
	fmt.Println(ch)
	fmt.Printf(`this is "\n"`)
	fmt.Printf("\n")

	//字符串可以通过直接+来拼接
	ch += "ade"
	fmt.Println(ch)

	//ch -= "lemon"但是不可以相减,这样的写法是错误的
	//fmt.Println(ch)
}
