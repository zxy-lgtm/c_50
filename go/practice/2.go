package main

import "fmt"

func main() {

	var score1, score2, tool int

	fmt.Scanf("%d %d", &score1, &score2)

	if score1 < 0 || score1 > 100 || score2 < 0 || score2 > 100 {
		tool = 0
	} else if score1 >= 60 && score2 >= 60 {
		tool = 1
	} else if score1 < 60 || score2 < 60 {
		tool = 2
	}

	switch tool {
	case 0:
		fmt.Println("it is error.")
	case 1:
		fmt.Println("it is pass.")
	case 2:
		fmt.Println("it is not pass.")
	}

}
