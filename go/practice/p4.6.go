package main

import (
	"fmt"
	"unicode/utf8"
)

var ch2 string

func main() {
	ch1 := "asSASA ddd dsjkdsjs dk"
	fmt.Println(len(ch1))
	fmt.Println(utf8.RuneCountInString(ch1))

	ch2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println(len(ch2))
	fmt.Println(utf8.RuneCountInString(ch2))
	//こん不是纯ascii码里的字符,一个占两个字节
}
