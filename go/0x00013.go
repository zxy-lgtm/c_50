package main

import "fmt"

type struct1 struct {
	a int
	b string
}

func main() {
	s := new(struct1)
	s.a = 1
	s.b = "w"
	fmt.Println(s.a, s.b, s)

	s2 := &struct1{2, "e"}
	fmt.Println(s2.a, s2)

}
