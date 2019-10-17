package main

import "fmt"

/*
005 最长回文子串 https://leetcode-cn.com/problems/longest-palindromic-substring

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
func main() {
	fmt.Print(longestPalindrome("aaabaaaa"))
}

// 执行用时 4 ms 内存消耗 2.2 MB
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 1
	for i := range s {
		for j := 1; i+j <= len(s)-1 && i-j >= 0; j++ {
			if (end-start)/2 > i || (end-start)/2 > len(s)-i || s[i+j] != s[i-j] {
				break
			}
			if 2*j+1 > end-start {
				start = i - j
				end = i + j + 1
			}
		}
	}
	for i := range s {
		for j := 1; i+j <= len(s)-1 && i-j+1 >= 0; j++ {
			if (end-start)/2 > i || (end-start)/2 > len(s)-i || s[i+j] != s[i-j+1] {
				break
			}
			if 2*j > end-start {
				start = i - j + 1
				end = i + j + 1
			}
		}
	}
	return string(s[start:end])
}
