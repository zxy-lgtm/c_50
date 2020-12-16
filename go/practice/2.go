package main

import (
	"fmt"
	"strconv"
)

func main() {

	mymap := make(map[int]string)
	mymap[0] = "ling"
	mymap[1] = "yi"
	mymap[2] = "er"
	mymap[3] = "san"
	mymap[4] = "si"
	mymap[5] = "wu"
	mymap[6] = "liu"
	mymap[7] = "qi"
	mymap[8] = "ba"
	mymap[9] = "jiu"

	var n string
	sum := 0
	fmt.Scan(&n)
	//fmt.Println(n)
	str := n[:]
	for _, ok := range str {
		//fmt.Println(key, (ok - '0'))
		numer := int(ok - '0')
		sum += numer
		//fmt.Println(sum)
	}

	newSum := strconv.Itoa(sum)
	for key, ok := range newSum[:] {
		number := int(ok - '0')
		if key == 0 {
			fmt.Printf("%s", mymap[number])
		} else {
			fmt.Printf(" %s", mymap[number])
		}
	}
}
