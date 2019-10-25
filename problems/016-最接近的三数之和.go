package main

import (
	"fmt"
)

/*
016 最接近的三数之和 https://leetcode-cn.com/problems/3sum-closest/

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/
func main() {
	fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
}

// 执行用时 12 ms 内存消耗 2.7 MB
func threeSumClosest(nums []int, target int) int {
	min := 999
	number := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if diff(nums[i]+nums[j]+nums[k], target) == 0 {
					return target
				}
				if diff(nums[i]+nums[j]+nums[k], target) < min {
					min = diff(nums[i]+nums[j]+nums[k], target)
					number = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return number
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
