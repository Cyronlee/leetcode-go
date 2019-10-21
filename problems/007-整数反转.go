package main

import (
	"fmt"
	"math"
)

/*
007 整数反转 https://leetcode-cn.com/problems/reverse-integer/

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
func main() {
	fmt.Print(reverse(1534236469))
}

// 执行用时 4 ms 内存消耗 2.2 MB
func reverse(x int) int {
	y := 0
	var r int
	for {
		if x == 0 {
			break
		}
		r = x % 10
		x = x / 10
		y = 10*y + r
	}
	if y > math.MaxInt32 || -y > math.MaxInt32 {
		return 0
	}
	return y
}
