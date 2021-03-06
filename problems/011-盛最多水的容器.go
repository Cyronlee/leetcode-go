package main

import (
	"fmt"
)

/*
011 盛最多水的容器 https://leetcode-cn.com/problems/container-with-most-water/

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg
图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。


示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// 执行用时 8 ms 内存消耗 5.7 MB
func maxArea(height []int) int {
	area := 0
	l := 0
	r := len(height) - 1
	for l < r {
		area = max(area, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

// 暴力法执行用时 16 ms 内存消耗 2.2 MB
func maxArea2(height []int) int {
	max := 0
	area := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				area = (j - i) * height[j]
			} else {
				area = (j - i) * height[i]
			}
			if area > max {
				max = area
			}
		}
	}
	return max
}
