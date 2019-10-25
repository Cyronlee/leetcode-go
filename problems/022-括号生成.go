package main

import "fmt"

/*
022 括号生成 https://leetcode-cn.com/problems/generate-parentheses/

给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func main() {
	fmt.Println(generateParenthesis(3))
}

// 执行用时 20 ms 内存消耗 7.7 MB
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	pre := generateParenthesis(n - 1)
	var temp []string
	for i := 0; i < len(pre); i++ {
		for j := 0; j < len(pre[i])/2+1; j++ {
			temp = appendDistinct(temp, pre[i][:j]+"()"+pre[i][j:])
		}
	}
	return temp
}

func appendDistinct(strs []string, str string) []string {
	for _, v := range strs {
		if v == str {
			return strs
		}
	}
	return append(strs, str)
}
