package main

import (
	"fmt"
	"strconv"
)

//我自己写的函数,不是很好(内存和运行速度),方法是将数字变为字符串再进行比较
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		s := strconv.Itoa(x)
		slice := s[:]
		//测试需要fmt.Println(slice)
		end := len(slice) - 1
		head := 0
		for end != head && end >= 0 {
			//测试需要fmt.Println(slice[end], slice[head])
			if slice[head] != slice[end] {
				return false
			}
			end--
			head++
		}
		return true
	}
}

//利用for-range来写
func isPalindrome2(x int) bool {
	var str = strconv.Itoa(x)
	if x < 0 {
		return false
	}
	for i := range str {
		if str[len(str)-1-i] != str[i] {
			return false
		}
	}
	return true
}

//数字反转比较
func isPalindrome3(x int) bool {
	// 倒序后  判断是不是和原来的数字相等
	if x < 0 {
		return false
	}
	var number int
	origin := x
	for x != 0 {
		number = number*10 + x%10
		if !(-(1<<31) <= number && number <= (1<<31)-1) {
			return false
		}
		x /= 10
	}
	return number == origin
}

func main() {
	//第一次提交没有通过的数字
	x := 100021
	fmt.Println(isPalindrome(x), isPalindrome2(x), isPalindrome3(x))
}
