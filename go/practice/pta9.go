package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputReader *bufio.Reader
	//var res []string
	inputReader = bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	if err != nil {
		return
	}
	s := strings.Fields(str)
	number := len(s) - 1
	count := 0
	if number%2 == 0 {
		count = number / 2
	} else {
		count = (number - 1) / 2
	}
	//fmt.Println(number)
	for key, value := range s {
		if key < count+1 {
			//fmt.Println(s[key], value, count)
			s[key] = s[number-key]
			s[number-key] = value
		} else {

		}
	}
	/*for key, value := range s {
		if key < number {
			res = append(res, value)
			res = append(res, ` `)
		} else {
			res = append(res, value)
		}
	}*/

	t := strings.Join(s, " ")
	fmt.Println(t)

}
