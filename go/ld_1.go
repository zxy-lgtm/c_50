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

func main() {
	fmt.Println("两数之和:\n下标为:")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
