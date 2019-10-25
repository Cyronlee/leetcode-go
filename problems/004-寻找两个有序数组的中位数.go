package main

import "fmt"

/*
004 寻找两个有序数组的中位数 https://leetcode-cn.com/problems/median-of-two-sorted-arrays

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

// 执行用时 20 ms 内存消耗 5.5 MB
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	x2 := length / 2
	x1 := x2
	if length%2 == 0 {
		x1 = x2 - 1
	}

	var i1, i2 = 0, 0
	var mid float64
	var mid1, mid2 float64
	for i := 0; i <= x2; i++ {
		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] < nums2[i2] {
				mid = float64(nums1[i1])
				i1++
			} else {
				mid = float64(nums2[i2])
				i2++
			}
		} else if i1 < len(nums1) {
			mid = float64(nums1[i1])
			i1++
		} else if i2 < len(nums2) {
			mid = float64(nums2[i2])
			i2++
		}
		if i == x1 {
			mid1 = mid
		}
		if i == x2 {
			mid2 = mid
		}
	}
	return (mid1 + mid2) / 2
}
