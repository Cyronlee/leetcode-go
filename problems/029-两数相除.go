package main

import (
	"fmt"
	"math"
)

/*
029 两数相除 https://leetcode-cn.com/problems/divide-two-integers/

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/
func main() {
	fmt.Println(divide(10, 2))
}

// 执行用时 4 ms 内存消耗 2.4 MB
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == -math.MaxInt32-1 && divisor == -1 {
		return math.MaxInt32
	}
	negative := false
	if dividend < 0 {
		negative = !negative
		dividend = - dividend
	}
	if divisor < 0 {
		negative = !negative
		divisor = -divisor
	}
	result := 0
	for i := 31; i >= 0; i-- {
		if (dividend >> uint(i)) >= divisor {
			result += 1 << uint(i)
			dividend -= divisor << uint(i)
		}
	}
	if negative {
		return -result
	}
	return result
}
