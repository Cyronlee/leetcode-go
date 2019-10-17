package main

/*
026 删除排序数组中的重复项 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/submissions/

给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
*/
func main() {
	nums := []int{1, 2, 3, 3}
	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	len := removeDuplicates(nums);
	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
	for i := 0; i < len; i++ {
		print(nums[i]);
	}
}

// 执行用时 60 ms 内存消耗 7.5 MB
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
