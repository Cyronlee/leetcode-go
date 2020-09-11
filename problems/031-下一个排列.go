package main

import "fmt"

/*
031 下一个排列 https://leetcode-cn.com/problems/next-permutation/

实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func main() {
	arr := []int{1, 2, 3}
	nextPermutation(arr)
	fmt.Println(arr)
}

// 执行用时 4 ms 内存消耗 2.6 MB
func nextPermutation(nums []int) {
	/*if len(nums) == 1 {
		return
	}*/
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j >= 0; j-- {
				if nums[j] > nums[i] {
					temp := nums[j]
					nums[j] = nums[i]
					nums[i] = temp
					break
				}
			}
			reverseBetween(nums, i+1, len(nums)-1)
			return
		}
	}
	reverseBetween(nums, 0, len(nums)-1)
}

func reverseBetween(nums []int, start int, end int) {
	for l, r := start, end; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
