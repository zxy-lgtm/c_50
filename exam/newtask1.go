package main

import "fmt"

func Repalce(s string, target, replace byte) string {
	//首先将字符串转化为byte数组,然后一个一个对比
	slice := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if slice[i] == target {
			slice[i] = replace
		}
	}
	//因为规定了返回值是string类型,所以要进行一次转换
	result := string(slice)
	return result
}

func main() {
	s := "You are interesting ,but they are boring."
	s = Repalce(s, 'i', 'e')
	fmt.Println(s)

}
