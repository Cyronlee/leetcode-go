package main

import "fmt"

/*
003 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func main() {
	fmt.Print(lengthOfLongestSubstring("abav"))
	fmt.Print(lengthOfLongestSubstring2("abav"))
}

// 执行用时 8 ms 内存消耗 2.7 MB
func lengthOfLongestSubstring(s string) int {
	var bytes []byte
	max := 0
	for _, b := range []byte(s) {
		if len(bytes) == 0 {
			bytes = append(bytes, b)
		} else {
			for i, c := range bytes {
				if c == b {
					bytes = bytes[i+1:]
				}
			}
			bytes = append(bytes, b)
		}
		if len(bytes) > max {
			max = len(bytes)
		}
	}
	return max
}

// 执行用时 12 ms 内存消耗 3.1 MB
func lengthOfLongestSubstring2(s string) int {
	max := 0
	start := 0
	lastOccur := make(map[byte]int)
	for index, b := range []byte(s) {
		if i, ok := lastOccur[b]; ok {
			if i >= start {
				start = i + 1
			}
		}
		if index-start+1 > max {
			max = index - start + 1
		}

		lastOccur[b] = index
	}
	return max
}
