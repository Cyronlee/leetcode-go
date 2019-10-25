package main

import "fmt"

/*
030 串联所有单词的子串 https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
*/
func main() {
	fmt.Println(findSubstring("aaabbbc", []string{"a", "a", "b", "b", "c"}))
}

// 执行用时 348 ms 内存消耗 7.1 MB
func findSubstring(s string, words []string) []int {
	var indices []int
	if len(words) == 0 {
		return indices
	}
	for i := 0; i < len(s)-len(words[0])*len(words)+1; i++ {
		var temp = make([]string, len(words))
		copy(temp, words)
		if containsSubstring(s, i, temp) {
			indices = append(indices, i)
		}
	}
	return indices
}

func containsSubstring(s string, index int, words []string) bool {
	l := len(words)
	if l == 0 {
		return true
	}
	for i := 0; i < l; i++ {
		if words[i] == s[index:index+len(words[i])] {
			index = index + len(words[i])
			words = append(words[:i], words[i+1:]...)
			break
		}
	}
	if len(words) == l {
		return false
	}
	return containsSubstring(s, index, words)
}

func findSubstring2(s string, words []string) []int {
	var indices []int
	combinations := allCombination(words)
	for _, c := range combinations {
		for i := 0; i < len(s)-len(c)+1; i++ {
			if s[i:i+len(c)] == c {
				indices = append(indices, i)
			}
		}
	}

	return indices
}

func allCombination(words []string) []string {
	if len(words) == 1 {
		return words
	}
	var com []string
	for _, word := range words {
		if len(com) == 0 {
			com = append(com, word)
			continue
		}
		var temp []string
		for i := 0; i < len(com); i++ {
			for j := 0; j < len(com[i])+1; j += len(word) {
				s := com[i][:j] + word + com[i][j:]
				contains := false
				for _, v := range temp {
					if v == s {
						contains = true
					}
				}
				if !contains {
					temp = append(temp, s)
				}
			}
		}
		com = temp
	}
	return com
}

/*c

cb bc

acb cab cba abc bac bca*/
