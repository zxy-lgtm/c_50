package main

import (
	"fmt"
	"log"
	"unicode"
)

type Rope string

func main() {
	var a Rope
	a = "qwertyuio1234567890"
	ch := a
	b := 'q'
	var ch byte = '\x42'
	log.Println(a, b, ch)
	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3)
	fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3)
	fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3)
	fmt.Printf("%U - %U - %U\n", ch1, ch2, ch3)
	if unicode.Letter(ch) {
		fmt.Println(ch)
	}
}
