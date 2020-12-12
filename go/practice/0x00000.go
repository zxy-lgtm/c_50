package main

import (
	"fmt"
	"log"
)

const (
	a = iota
	b
	c
	d
)

const (
	_ = iota
	mon
	tue
	wen
	thur
	fri
	st
	sun
)

const (
	e = b << iota
	f
	g
	h
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(mon, tue, wen, thur, fri, st, sun)
	log.Println(e, f, g, h)

}
