package main

import "fmt"

func Repalce(s string, target, replace byte) string {
	//首先将字符串转化为byte数组,然后一个一个对比
	slice := []byte(s)
	if len(slice) == 0 {
		return ""
	}
	for i, value := range slice {
		//fmt.Println(value)
		if value == target {
			slice[i] = replace
			//fmt.Println("!!!")
		}
	}
	//因为规定了返回值是string类型,所以要进行一次转换
	result := string(slice)
	return result
}

func main() {
	s := "You are interesting."
	s = Repalce(s, 'i', 'e')
	s1 := ""
	s1 = Repalce(s1, 'i', 'e')
	fmt.Println(s, s1)

}
