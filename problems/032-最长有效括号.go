package main

import "fmt"

/*
032 最长有效括号 https://leetcode-cn.com/problems/longest-valid-parentheses

给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"

示例 2:
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

// 执行用时 0 ms 内存消耗 3.5 MB
func longestValidParentheses(s string) int {
	stack := []int{-1}
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else if l := i - stack[len(stack)-1]; l > max {
				max = l
			}
		}
	}
	return max
}
