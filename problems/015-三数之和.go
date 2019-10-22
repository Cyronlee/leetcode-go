package main

import (
	"fmt"
	"sort"
)

/*
015 三数之和 https://leetcode-cn.com/problems/3sum/

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func main() {
	fmt.Print(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}

// 执行用时 968 ms 内存消耗 205 MB
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// fmt.Println(nums)
	var arrays [][]int
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for r > l {
			if nums[i]+nums[l]+nums[r] == 0 {
				if !contains3(arrays, nums[i], nums[l], nums[r]) {
					arrays = append(arrays, []int{nums[i], nums[l], nums[r]})
				}
			}
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return arrays
}

func contains3(arrays [][]int, a int, b int, c int) bool {
	for i := 0; i < len(arrays); i++ {
		if arrays[i][0] == a &&
			arrays[i][1] == b &&
			arrays[i][2] == c {
			return true
		}
	}
	return false
}
