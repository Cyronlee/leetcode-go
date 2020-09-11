package main

/*
027 移除元素 https://leetcode-cn.com/problems/remove-element/

*/
func main() {
	nums := []int{3, 2, 2, 3}
	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	len := removeElement(nums, 3)
	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
	for i := 0; i < len; i++ {
		print(nums[i])
	}
}

// 执行用时 0 ms 内存消耗 2.4 MB
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l := len(nums)
	for {
		move := false
		for i := 0; i < l-1; i++ {
			if nums[i] == val {
				move = true
			}
			if move {
				nums[i] = nums[i+1]
			}
		}
		if !move {
			break
		}
		l--
	}
	if nums[l-1] == val {
		l--
	}
	return l
}
