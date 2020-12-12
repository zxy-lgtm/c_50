package main

import (
	"fmt"
)

func main() {

	s := make([]byte, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = s[2:4]
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(len(s2), cap(s2))
	//s2[1] == 't'
	s2[1] = 't'
	fmt.Println(len(s2), cap(s2))

}
