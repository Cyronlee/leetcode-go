package main

import (
	"fmt"
	"sort"
)

/*
018 四数之和 https://leetcode-cn.com/problems/4sum/

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func main() {
	fmt.Println(fourSum([]int{-1, 4, -4, 3, -2, 3, 0}, -7))
}

// 执行用时 12 ms 内存消耗 3.1 MB
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	//fmt.Println(nums)
	var arrays [][]int
	for i := 0; i < len(nums)-3; i++ {
		for j := len(nums) - 1; j > 2; j-- {
			l := i + 1
			r := j - 1
			for r > l {
				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					if !contains4(arrays, nums[i], nums[j], nums[l], nums[r]) {
						arrays = append(arrays, []int{nums[i], nums[j], nums[l], nums[r]})
					}
				}
				if nums[i]+nums[j]+nums[l]+nums[r] > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return arrays
}

func contains4(arrays [][]int, a int, b int, c int, d int) bool {
	for i := 0; i < len(arrays); i++ {
		if arrays[i][0] == a &&
			arrays[i][1] == b &&
			arrays[i][2] == c &&
			arrays[i][3] == d {
			return true
		}
	}
	return false
}
