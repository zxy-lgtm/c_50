package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice = slice[0 : i+1]
		slice[i] = i
	}
	fmt.Println(slice, len(slice[1:1]), len(slice[1:2]))

	sl1 := []int{1, 2, 3}
	sl2 := make([]int, 10)
	fmt.Println(sl1, sl2)
	n := copy(sl2, sl1)
	sl2 = append(sl2[:3], 4, 5, 6)
	sl2 = append(sl2[:6], 7, 8, 9)
	sl2 = append(sl2[:9], 10, 11, 12)
	fmt.Println(sl1, sl2, n)

	s := "12344sdffrr"
	c := []byte(s)
	fmt.Println(c)
	for _, ch := range c {
		ch--
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	d := "\u00ff\u754c"
	for i, c := range d {
		fmt.Printf("%d:%c ", i, c)
	}
	fmt.Println()

	//修改字符串
	sli := []byte(s) //转化为字节数组
	sli[2] = 'w'
	ch1 := string(sli) //转化为字符串
	fmt.Println(ch1)

	ch := "1234567"
	ch2 := "1234567"
	ch3 := "12344567"
	if Compare(ch, ch2) {
		fmt.Println("true")
	}
	if Compare(ch2, ch3) {
		fmt.Println("true")
	}

	ch6 := [...]int{1, 2, 3, 2, 5, 8, 6, 7, 4, 33, 2}
	//ch5 := make([]int, 11)
	ch5 := ch6[:5]
	fmt.Println(ch6)
	sort.Ints(ch5)
	fmt.Println(ch5, ch6)
	ch5 = append(ch6[:2], ch6[3:]...)
	fmt.Println(ch5)
	ch7 := ch5[:]
	ch7 = append(ch5, make([]int, 3)...)
	fmt.Println(ch7)

	a := []int{1, 2, 3, 4, 5}
	a1 := a[:]
	a1 = append(a[:2], append([]int{34}, a[2:]...)...)
	fmt.Println(a1)
}

//比较字符串
func Compare(ch1, ch2 string) bool {
	for i := 0; i < len(ch1) && i < len(ch2); i++ {
		if ch1[i] != ch2[i] {
			return false
		}
	}

	if len(ch1) != len(ch2) {
		return false
	}
	return true
}

/*
删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)

切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)

为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)

在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)

在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)

在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)

取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]

将元素 x 追加到切片 a：a = append(a, x)
*/
