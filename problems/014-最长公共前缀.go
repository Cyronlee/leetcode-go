package main

import (
	"fmt"
)

/*
014 最长公共前缀 https://leetcode-cn.com/problems/longest-common-prefix/submissions/

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
func main() {
	fmt.Print(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

// 执行用时 0 ms 内存消耗 2.4 MB
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix []byte
	x := 0
	end := false
	var char byte
	for !end {
		if x >= len(strs[0]) {
			break
		}
		eq := true
		char = strs[0][x]
		for i := 0; i < len(strs); i++ {
			if x >= len(strs[i]) || strs[i][x] != char {
				eq = false
				end = true
				break
			}
		}
		if eq {
			prefix = append(prefix, strs[0][x])
		}
		x++
	}
	return fmt.Sprintf("%s", prefix)
}
