package main

import (
	"fmt"
)

/*
017 电话号码的字母组合 https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/
func main() {
	fmt.Print(letterCombinations("234"))
}

// 执行用时 0 ms 内存消耗 2.6 MB
func letterCombinations(digits string) []string {
	keys := make(map[int32][]string)
	keys['2'] = []string{"a", "b", "c"}
	keys['3'] = []string{"d", "e", "f"}
	keys['4'] = []string{"g", "h", "i"}
	keys['5'] = []string{"j", "k", "l"}
	keys['6'] = []string{"m", "n", "o"}
	keys['7'] = []string{"p", "q", "r", "s"}
	keys['8'] = []string{"t", "u", "v"}
	keys['9'] = []string{"w", "x", "y", "z"}
	var strs []string
	for _, digit := range digits {
		letters := keys[digit]
		if len(strs) == 0 {
			strs = letters
		} else {
			var temp []string
			for _, old := range strs {
				for _, letter := range letters {
					temp = append(temp, old+letter)
				}
			}
			strs = temp
		}
	}
	return strs
}
