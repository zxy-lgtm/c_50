package main

import "fmt"

//看了解析之后的写法
func romanToInt(s string) int {
	mymap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	left, sum := 0, 0
	//因为要比较左右所以从最右边开始,特殊情况就只有一种了
	for i := len(s) - 1; i >= 0; i-- {
		right := mymap[s[i]]
		//放在左边的话就是减,放在右边的话就是加
		if left > right {
			sum -= right
		} else {
			sum += right
		}
		left = right
	}

	return sum

}

//自己写的
func myromanToInt(s string) int {
	//首先将所有罗马数字对应的阿拉伯数字存再map里面
	romanMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	//sum是返回值
	sum := 0
	//将字符串转化为byte型数组切片,方便遍历
	romanBytes := []byte(s)

	for idx := 0; idx < len(romanBytes); idx++ {
		//这是为了算像III这种情况的罗马数字,因为II没有在map里面,它将不会执行下面那个if语句里的内容...
		if idx < len(romanBytes)-1 {
			/*因为map中为string类型,所以要进行一次转换,而因为map中的key最多两个byte,我们不妨直接一次性转化两个byte
			,下面这个判断是ok存在就执行*/
			if v, ok := romanMap[string(romanBytes[idx:idx+2])]; ok {
				sum += v
				idx += 2
				continue
			}
		}
		sum += romanMap[string(romanBytes[idx])]

	}
	return sum
}
func main() {
	s := "III"
	fmt.Println(romanToInt(s), myromanToInt(s))
}
