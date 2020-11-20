package main

import "fmt"

//暴力破解法
func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if val == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

//利用map来解题
func twoSum2(nums []int, target int) []int {
	slice := []int{}
	map1 := make(map[int]int)
	for i, k := range nums {
		if value, exist := map1[target-k]; exist {
			slice = append(slice, value)
			slice = append(slice, i)
		}
		map1[k] = i
	}
	return slice
}

func main() {
	fmt.Println("两数之和:\n下标为:")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}
