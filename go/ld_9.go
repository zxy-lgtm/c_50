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
		for end != head && end >= head {
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
	for i, _ := range str {
		//fmt.Println(i)
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

//官方题解,反转一半数字,这样就不用考虑整数溢出的情况
func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	//第一次提交没有通过的数字
	x := 100021
	fmt.Println(isPalindrome(x), isPalindrome2(x), isPalindrome3(x))
}
