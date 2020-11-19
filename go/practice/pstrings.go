package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "hello,world!"
	//判断字符串是否以hello开头,是否以hello结尾
	fmt.Printf("%t\n%t\n", strings.HasPrefix(myString, "hello"), strings.HasSuffix(myString, "hello"))
	//这里演示了\"\"的用法
	fmt.Println("\"--------这里是分割线----------\"")

	//判断字符串是否包含hello,ld,www,顺便这里演示了fmt.Printf可以换行不会报错
	fmt.Printf("%t %t %t \n", strings.Contains(myString, "hello"),
		strings.Contains(myString, "ld"), strings.Contains(myString, "www"))
	fmt.Println("\"--------这里是分割线----------\"")

	//寻找字符串第一次出现l的位置,最后一次出现的位置,如果没有则返回-1,注意返回的是下标!!!!!真实的第几个是下标加1
	fmt.Println(strings.Index(myString, "l"), strings.LastIndex(myString, "l"), strings.LastIndex(myString, "q"))

	fmt.Println("如果有非ascii字符的话~")
	myString2 := "这是我的字符串"
	//将汉字转化为int再寻找,但注意一个汉字占三个字节,所以除以汉字本身得到下标~
	fmt.Println(strings.IndexRune(myString2, '我') / len("我"))
	fmt.Println("\"--------这里是分割线----------\"")

	//将字符串中的前2个l字符全部替换成p,并且返回一个新字符串
	fmt.Println(strings.Replace(myString, "l", "p", 2))
	fmt.Println("\"--------这里是分割线----------\"")

	//统计字符串中l出现的次数
	fmt.Println(strings.Count(myString, "l"))
	fmt.Println("\"--------这里是分割线----------\"")

	//重复输出3次mystring里面的字符串
	fmt.Println(strings.Repeat(myString, 3))
	fmt.Println("\"--------这里是分割线----------\"")

	//将字符串大小写改变
	myString = strings.ToUpper(myString)
	fmt.Println(myString)
	myString = strings.ToLower(myString)
	fmt.Println(myString)
	fmt.Println("\"--------这里是分割线----------\"")

	//修剪字符串,包括空白符号,开头结尾的任意字符
	myString += "    "
	fmt.Println(myString, "#") //加上一个#来确定空白符的存在
	myString = strings.TrimSpace(myString)
	fmt.Println(myString, "#")
	//剪去从开头开始的第一个字符和从结尾开始的字符,不能分开
	fmt.Println(strings.Trim(myString, "hello"), strings.Trim(myString, "l"),
		strings.Trim(myString, "world!"))
	//剪去开头和结尾,可以分开
	fmt.Println(strings.TrimLeft(myString, "hello"), strings.TrimRight(myString, "world!"))
	fmt.Println("\"--------这里是分割线----------\"")

	//将字符串转成切片,Fields按照空白符号的多少来确定长度,Split按照自定义来确定长度
	slice := strings.Fields(myString)
	fmt.Println(slice, len(slice))
	myString3 := "h e l l o"
	fmt.Println(strings.Fields(myString3), len(strings.Fields(myString3)))
	slice2 := strings.Split(myString, ",")
	fmt.Println(slice2, len(slice2))
	fmt.Println("\"--------这里是分割线----------\"")

	//拼接slice和一个字符串变成一个新的字符串,注意是将这个字符串塞进这个切片的每一个断点
	fmt.Println(strings.Join(slice2, "!!!!"))
}
